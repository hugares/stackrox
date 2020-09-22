import React, { Component } from 'react';
import PropTypes from 'prop-types';
import * as Icon from 'react-feather';
import { connect } from 'react-redux';
import ReactRouterPropTypes from 'react-router-prop-types';
import { withRouter } from 'react-router-dom';
import { capitalize } from 'lodash';
import { createStructuredSelector } from 'reselect';

import { types as deploymentTypes } from 'reducers/deployments';
import { selectors } from 'reducers';
import { sortValue } from 'sorters/sorters';
import { filterModes, filterLabels } from 'constants/networkFilterModes';
import { getNetworkFlows } from 'utils/networkGraphUtils';
import Panel from 'Components/Panel';
import Loader from 'Components/Loader';
import TablePagination from 'Components/TablePagination';
import NoResultsMessage from 'Components/NoResultsMessage';
import Table, { rtTrActionsClassName } from 'Components/Table';
import RowActionButton from 'Components/RowActionButton';
import wizardStages from '../wizardStages';

class NamespaceDetails extends Component {
    static propTypes = {
        wizardOpen: PropTypes.bool.isRequired,
        wizardStage: PropTypes.string.isRequired,
        isFetchingNamespace: PropTypes.bool,
        onClose: PropTypes.func.isRequired,
        namespace: PropTypes.shape({
            id: PropTypes.string,
            deployments: PropTypes.arrayOf(PropTypes.shape({})),
        }),
        networkGraphRef: PropTypes.shape({
            setSelectedNode: PropTypes.func,
            selectedNode: PropTypes.shape({}),
            onNodeClick: PropTypes.func,
            getNodeData: PropTypes.func,
        }),
        filterState: PropTypes.number.isRequired,
        history: ReactRouterPropTypes.history.isRequired,
    };

    static defaultProps = {
        namespace: {
            id: null,
            deployments: [],
        },
        isFetchingNamespace: false,
        networkGraphRef: null,
    };

    constructor(props) {
        super(props);
        this.state = {
            page: 0,
            selectedNode: null,
        };
    }

    // TODO: refactor this component
    //   1. if leaving a class component, use `getDerivedStateFromProps` instead
    //   2. change to functional component with hooks
    // eslint-disable-next-line camelcase
    UNSAFE_componentWillReceiveProps = () => {
        this.setState({ selectedNode: null });
    };

    highlightNode = ({ data }) => {
        const { networkGraphRef } = this.props;
        if (data) {
            this.props.history.push(`/main/network/${data.deploymentId}`);
            networkGraphRef.setSelectedNode(data);
            this.setState({ selectedNode: data });
        }
    };

    navigate = ({ data }) => () => {
        const { onNodeClick } = this.props.networkGraphRef;
        if (data) {
            this.props.history.push(`/main/network/${data.deploymentId}`);
            onNodeClick(data);
        }
    };

    setTablePage = (newPage) => {
        this.setState({ page: newPage });
    };

    renderRowActionButtons = (node) => {
        return (
            <div className="border-2 border-r-2 border-base-400 bg-base-100 flex">
                <RowActionButton
                    text="Navigate to Deployment"
                    onClick={this.navigate(node)}
                    icon={<Icon.ArrowUpRight className="my-1 h-4 w-4" />}
                />
            </div>
        );
    };

    onPanelClose = () => {
        const { onClose, history } = this.props;
        history.push('/main/network');
        onClose();
    };

    renderTable() {
        const { namespace, filterState } = this.props;
        const filterStateString =
            filterState !== filterModes.all ? capitalize(filterLabels[filterState]) : 'Network';

        const columns = [
            {
                Header: 'Deployment',
                accessor: 'data.name',
                Cell: ({ value }) => <span>{value}</span>,
            },
            {
                Header: `${filterStateString} Flows`,
                accessor: 'data.edges',
                Cell: ({ value }) => {
                    const { networkFlows } = getNetworkFlows(value, filterState);
                    return <span>{networkFlows.length}</span>;
                },
                sortMethod: sortValue,
            },
            {
                accessor: '',
                headerClassName: 'hidden',
                className: rtTrActionsClassName,
                Cell: ({ original }) => this.renderRowActionButtons(original),
            },
        ];
        const rows = namespace.deployments;
        if (!rows.length) {
            return <NoResultsMessage message="No namespace deployments" />;
        }
        return (
            <Table
                rows={rows}
                columns={columns}
                onRowClick={this.highlightNode}
                noDataText="No namespace deployments"
                page={this.state.page}
                idAttribute="data.id"
                selectedRowId={this.state.selectedNode && this.state.selectedNode.id}
            />
        );
    }

    render() {
        const { namespace, wizardOpen, wizardStage, isFetchingNamespace } = this.props;
        if (!namespace || !wizardOpen || wizardStage !== wizardStages.namespaceDetails) {
            return null;
        }
        const paginationComponent = (
            <TablePagination
                page={this.state.page}
                dataLength={namespace && namespace.deployments && namespace.deployments.length}
                setPage={this.setTablePage}
            />
        );
        const subHeaderText = `${namespace.deployments.length} Deployment${
            namespace.deployments.length === 1 ? '' : 's'
        }`;
        const content = isFetchingNamespace ? <Loader /> : <div>{this.renderTable()}</div>;

        return (
            <Panel header={namespace.id} onClose={this.onPanelClose}>
                <Panel
                    header={subHeaderText}
                    headerComponents={paginationComponent}
                    isUpperCase={false}
                    className="w-full h-full bg-base-100"
                >
                    <div className="w-full h-full">{content}</div>
                </Panel>
            </Panel>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    wizardOpen: selectors.getNetworkWizardOpen,
    wizardStage: selectors.getNetworkWizardStage,
    namespace: selectors.getSelectedNamespace,
    isFetchingNamespace: (state) =>
        selectors.getLoadingStatus(state, deploymentTypes.FETCH_DEPLOYMENTS),
    networkGraphRef: selectors.getNetworkGraphRef,
    filterState: selectors.getNetworkGraphFilterMode,
});

export default withRouter(connect(mapStateToProps)(NamespaceDetails));
