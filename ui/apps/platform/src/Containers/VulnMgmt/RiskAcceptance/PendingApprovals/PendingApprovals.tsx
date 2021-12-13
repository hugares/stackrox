import React, { ReactElement } from 'react';
import { Bullseye, Spinner } from '@patternfly/react-core';
import { useQuery, useApolloClient } from '@apollo/client';

import {
    GetVulnerabilityRequestsData,
    GetVulnerabilityRequestsVars,
    GET_VULNERABILITY_REQUESTS,
} from '../vulnerabilityRequests.graphql';

import PendingApprovalsTable from './PendingApprovalsTable';

function PendingApprovals(): ReactElement {
    const client = useApolloClient();
    const { loading: isLoading, data } = useQuery<
        GetVulnerabilityRequestsData,
        GetVulnerabilityRequestsVars
    >(GET_VULNERABILITY_REQUESTS, {
        variables: {
            query: 'Request Status:PENDING,APPROVED_PENDING_UPDATE+Expired Request:false',
            pagination: {
                limit: 20,
                offset: 0,
                sortOption: {
                    field: 'id',
                    reversed: false,
                },
            },
        },
        fetchPolicy: 'network-only',
    });

    async function updateTable() {
        await client.refetchQueries({
            include: [GET_VULNERABILITY_REQUESTS],
        });
    }

    if (isLoading) {
        return (
            <Bullseye>
                <Spinner size="sm" />
            </Bullseye>
        );
    }

    const rows = data?.results || [];

    return <PendingApprovalsTable rows={rows} updateTable={updateTable} isLoading={isLoading} />;
}

export default PendingApprovals;
