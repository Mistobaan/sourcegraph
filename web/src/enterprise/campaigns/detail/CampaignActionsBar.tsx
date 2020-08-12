import React from 'react'
import { Link } from '../../../../../shared/src/components/Link'
import { CampaignFields } from '../../../graphql-operations'

interface Props {
    campaign: Pick<CampaignFields, 'name' | 'closedAt' | 'viewerCanAdminister' | 'namespace' | 'changesets'>
}

export const CampaignActionsBar: React.FunctionComponent<Props> = ({ campaign }) => {
    const campaignClosed = !!campaign.closedAt

    const percentComplete = (campaign.changesets.stats.total === 0
        ? 0
        : ((campaign.changesets.stats.closed + campaign.changesets.stats.merged) / campaign.changesets.stats.total) *
          100
    ).toFixed(0)

    return (
        <>
            <div className="mb-2">
                <span>
                    <Link to="/campaigns">Campaigns</Link>
                </span>
                <span className="text-muted d-inline-block mx-1">/</span>
                <span>
                    <Link to={campaign.namespace.url}>{campaign.namespace.namespaceName}</Link>
                </span>
                <span className="text-muted d-inline-block mx-1">/</span>
                <span>{campaign.name}</span>
            </div>
            <div className="d-flex mb-2 position-relative">
                <div>
                    <h1 className="mb-1">{campaign.name}</h1>
                    <div>
                        <CampaignStateBadge isClosed={campaignClosed} />
                        <small className="text-muted">
                            {percentComplete}% complete. {campaign.changesets.totalCount} changesets total
                        </small>
                    </div>
                </div>
            </div>
        </>
    )
}

export const CampaignStateBadge: React.FunctionComponent<{ isClosed: boolean }> = ({ isClosed }) => {
    if (isClosed) {
        return <span className="badge badge-danger text-uppercase">Closed</span>
    }
    return <span className="badge badge-success text-uppercase">Open</span>
}
