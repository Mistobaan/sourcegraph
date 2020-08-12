import React from 'react'
import ProgressCheckIcon from 'mdi-react/ProgressCheckIcon'
import SourceMergeIcon from 'mdi-react/SourceMergeIcon'
import SourceBranchIcon from 'mdi-react/SourceBranchIcon'
import CheckCircleOutlineIcon from 'mdi-react/CheckCircleOutlineIcon'
import classNames from 'classnames'
import { CampaignFields } from '../../../graphql-operations'
import { CampaignStateBadge } from './CampaignActionsBar'
import SourcePullIcon from 'mdi-react/SourcePullIcon'

interface CampaignStatsCardProps extends Pick<CampaignFields['changesets'], 'stats'> {
    closedAt: CampaignFields['closedAt']
}

export const CampaignStatsCard: React.FunctionComponent<CampaignStatsCardProps> = ({ stats, closedAt }) => {
    const percentComplete = stats.total === 0 ? 0 : (((stats.closed + stats.merged) / stats.total) * 100).toFixed(0)
    const isCompleted = stats.closed + stats.merged === stats.total
    let CampaignStatusIcon = ProgressCheckIcon
    if (isCompleted) {
        CampaignStatusIcon = CheckCircleOutlineIcon
    }
    return (
        <div className="card">
            <div className="card-body p-3">
                <div className="d-flex flex-wrap justify-content-between align-items-center">
                    <div className="d-flex align-items-center flex-grow-1">
                        <h2 className="m-0 mr-3">
                            <CampaignStateBadge isClosed={!!closedAt} />
                        </h2>
                        <h1 className="d-inline mb-0">
                            <CampaignStatusIcon
                                className={classNames(
                                    'icon-inline mr-2',
                                    isCompleted && 'text-success',
                                    !isCompleted && 'text-muted'
                                )}
                            />
                        </h1>{' '}
                        <span className="lead">{percentComplete}% complete</span>
                    </div>
                    <CampaignStatsTotalAction count={stats.total} />
                    <CampaignStatsUnpublishedAction count={stats.unpublished} />
                    <CampaignStatsOpenAction count={stats.open} />
                    <CampaignStatsClosedAction count={stats.closed} />
                    <CampaignStatsMergedAction count={stats.merged} />
                </div>
            </div>
        </div>
    )
}

export const CampaignStatsTotalAction: React.FunctionComponent<{ count: number }> = ({ count }) => (
    <div className="m-0 mr-3 flex-grow-0 flex-shrink-0 text-nowrap d-flex flex-column align-items-center justify-content-center">
        <span className="campaign-stats-card__changesets-pill">
            <span className="badge badge-pill badge-secondary">{count}</span>
        </span>
        <span className="text-muted">changesets</span>
    </div>
)

export const CampaignStatsUnpublishedAction: React.FunctionComponent<{ count: number }> = ({ count }) => (
    <div className="m-0 mx-3 flex-grow-0 flex-shrink-0 text-nowrap d-flex flex-column align-items-center justify-content-center">
        <SourceBranchIcon className="text-muted" />
        <span className="text-muted">{count} unpublished</span>
    </div>
)
export const CampaignStatsOpenAction: React.FunctionComponent<{ count: number }> = ({ count }) => (
    <div className="m-0 mx-3 flex-grow-0 flex-shrink-0 text-nowrap d-flex flex-column align-items-center justify-content-center">
        <SourcePullIcon className="text-success" />
        <span className="text-muted">{count} open</span>
    </div>
)
export const CampaignStatsClosedAction: React.FunctionComponent<{ count: number }> = ({ count }) => (
    <div className="m-0 mx-3 flex-grow-0 flex-shrink-0 text-nowrap d-flex flex-column align-items-center justify-content-center">
        <SourcePullIcon className="text-danger" />
        <span className="text-muted">{count} closed</span>
    </div>
)
export const CampaignStatsMergedAction: React.FunctionComponent<{ count: number }> = ({ count }) => (
    <div className="m-0 ml-3 flex-grow-0 flex-shrink-0 text-nowrap d-flex flex-column align-items-center justify-content-center">
        <SourceMergeIcon className="text-merged" />
        <span className="text-muted">{count} merged</span>
    </div>
)
