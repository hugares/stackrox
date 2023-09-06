import React, { ReactElement } from 'react';
import {
    CheckCircleIcon,
    DownloadIcon,
    ExclamationCircleIcon,
    ExclamationTriangleIcon,
    HelpIcon,
    InProgressIcon,
    PendingIcon,
} from '@patternfly/react-icons';
import { Button, Flex, FlexItem, Tooltip } from '@patternfly/react-core';

import { ReportSnapshot } from 'services/ReportsService.types';

export type ReportJobStatusProps = {
    reportSnapshot: ReportSnapshot;
    areDownloadActionsDisabled: boolean;
    onDownload: () => void;
};

const genericMsg =
    'An issue was encountered. Please try again later. If the issue persists, please contact support.';

function ReportJobStatus({
    reportSnapshot,
    areDownloadActionsDisabled,
    onDownload,
}: ReportJobStatusProps): ReactElement {
    const { reportStatus, isDownloadAvailable } = reportSnapshot;

    let statusColorClass = '';
    let statusIcon: ReactElement;
    let statusText: ReactElement;

    if (
        reportStatus.runState === 'GENERATED' &&
        isDownloadAvailable &&
        areDownloadActionsDisabled
    ) {
        statusColorClass = 'pf-u-disabled-color-100';
        statusIcon = <DownloadIcon title="Report download was successfully prepared" />;
        statusText = (
            <Flex
                direction={{ default: 'row' }}
                spaceItems={{ default: 'spaceItemsSm' }}
                alignItems={{ default: 'alignItemsCenter' }}
            >
                <FlexItem>
                    <p>Ready for download</p>
                </FlexItem>
                <FlexItem>
                    <Tooltip
                        content={
                            <div>
                                Only the requestor of the download has the authority to access or
                                remove it.
                            </div>
                        }
                    >
                        <HelpIcon />
                    </Tooltip>
                </FlexItem>
            </Flex>
        );
    } else if (
        reportStatus.runState === 'GENERATED' &&
        isDownloadAvailable &&
        !areDownloadActionsDisabled
    ) {
        statusColorClass = 'pf-u-primary-color-100';
        statusIcon = <DownloadIcon title="Report download was successfully prepared" />;
        statusText = (
            <Button variant="link" isInline className={statusColorClass} onClick={onDownload}>
                Ready for download
            </Button>
        );
    } else if (reportStatus.runState === 'GENERATED' && !isDownloadAvailable) {
        statusColorClass = 'pf-u-disabled-color-100';
        statusIcon = <DownloadIcon title="Report download was deleted" />;
        statusText = (
            <Flex
                direction={{ default: 'row' }}
                spaceItems={{ default: 'spaceItemsSm' }}
                alignItems={{ default: 'alignItemsCenter' }}
            >
                <FlexItem>
                    <p>Download deleted</p>
                </FlexItem>
                <FlexItem>
                    <Tooltip
                        content={
                            <div>
                                The download was deleted. Please generate a new download if needed.
                            </div>
                        }
                    >
                        <HelpIcon />
                    </Tooltip>
                </FlexItem>
            </Flex>
        );
    } else if (reportStatus.runState === 'DELIVERED') {
        statusColorClass = 'pf-u-success-color-100';
        statusIcon = <CheckCircleIcon title="Report was successfully delivered" />;
        statusText = <p className="pf-u-success-color-100">Successfully delivered</p>;
    } else if (reportStatus.runState === 'FAILURE') {
        statusColorClass = 'pf-u-danger-color-100';
        statusIcon = (
            <Tooltip content={reportStatus?.errorMsg || genericMsg}>
                <ExclamationCircleIcon title="Report run was unsuccessful" />
            </Tooltip>
        );
        statusText = <p>Error</p>;
    } else if (reportStatus.runState === 'PREPARING') {
        statusIcon = <InProgressIcon title="Report run is preparing" />;
        statusText = <p>Preparing</p>;
    } else if (reportStatus.runState === 'WAITING') {
        statusIcon = <PendingIcon title="Report run is waiting" />;
        statusText = <p>Waiting</p>;
    } else {
        statusColorClass = 'pf-u-warning-color-100';
        statusIcon = (
            <Tooltip content="Please contact support for more help.">
                <ExclamationTriangleIcon title="Report run status is unknown" />
            </Tooltip>
        );
        statusText = <p>Unknown status</p>;
    }

    return (
        <Flex alignItems={{ default: 'alignItemsCenter' }} className={statusColorClass}>
            <FlexItem>{statusIcon}</FlexItem>
            <FlexItem>{statusText}</FlexItem>
        </Flex>
    );
}

export default ReportJobStatus;
