import { LoadingSpinner } from '@sourcegraph/react-loading-spinner'
import AlertCircleIcon from 'mdi-react/AlertCircleIcon'
import React, { useEffect, useMemo } from 'react'
import { HeroPage } from '../../../components/HeroPage'
import { PageTitle } from '../../../components/PageTitle'
import { isEqual } from 'lodash'
import { fetchCampaignById } from './backend'
import { useObservable } from '../../../../../shared/src/util/useObservable'
import * as H from 'history'
import { CampaignBurndownChart } from './BurndownChart'
import { Subject, of, merge } from 'rxjs'
import { switchMap, distinctUntilChanged } from 'rxjs/operators'
import { ThemeProps } from '../../../../../shared/src/theme'
import { CampaignActionsBar } from './CampaignActionsBar'
import { CampaignChangesets } from './changesets/CampaignChangesets'
import { ExtensionsControllerProps } from '../../../../../shared/src/extensions/controller'
import { PlatformContextProps } from '../../../../../shared/src/platform/context'
import { TelemetryProps } from '../../../../../shared/src/telemetry/telemetryService'
import { CampaignFields, Scalars } from '../../../graphql-operations'
import { CampaignInfoCard } from './CampaignInfoCard'

interface Props extends ThemeProps, ExtensionsControllerProps, PlatformContextProps, TelemetryProps {
    /**
     * The campaign ID.
     */
    campaignID: Scalars['ID']
    history: H.History
    location: H.Location

    /** For testing only. */
    _fetchCampaignById?: typeof fetchCampaignById
}

/**
 * The area for a single campaign.
 */
export const CampaignDetails: React.FunctionComponent<Props> = ({
    campaignID,
    history,
    location,
    isLightTheme,
    extensionsController,
    platformContext,
    telemetryService,
    _fetchCampaignById = fetchCampaignById,
}) => {
    /** Retrigger fetching */
    const campaignUpdates = useMemo(() => new Subject<void>(), [])

    useEffect(() => {
        telemetryService.logViewEvent(campaignID ? 'CampaignDetailsPage' : 'NewCampaignPage')
    }, [campaignID, telemetryService])

    const campaign: CampaignFields | null | undefined = useObservable(
        useMemo(
            () =>
                merge(of(undefined), campaignUpdates).pipe(
                    switchMap(() => _fetchCampaignById(campaignID)),
                    distinctUntilChanged((a, b) => isEqual(a, b))
                ),
            [campaignID, campaignUpdates, _fetchCampaignById]
        )
    )

    // Is loading.
    if (campaign === undefined) {
        return (
            <div className="text-center">
                <LoadingSpinner className="icon-inline mx-auto my-4" />
            </div>
        )
    }
    // Campaign was not found
    if (campaign === null) {
        return <HeroPage icon={AlertCircleIcon} title="Campaign not found" />
    }

    return (
        <>
            <PageTitle title={campaign.name} />
            <CampaignActionsBar campaign={campaign} />
            <CampaignInfoCard
                history={history}
                author={campaign.initialApplier}
                createdAt={campaign.createdAt}
                description={campaign.description}
            />
            <h3 className="mt-4 mb-2">Progress</h3>
            <CampaignBurndownChart changesetCountsOverTime={campaign.changesetCountsOverTime} history={history} />
            <CampaignChangesets
                campaignID={campaign.id}
                viewerCanAdminister={campaign.viewerCanAdminister}
                changesetUpdates={campaignUpdates}
                campaignUpdates={campaignUpdates}
                history={history}
                location={location}
                isLightTheme={isLightTheme}
                extensionsController={extensionsController}
                platformContext={platformContext}
                telemetryService={telemetryService}
            />
        </>
    )
}
