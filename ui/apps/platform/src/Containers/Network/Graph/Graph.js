import React, { Component } from 'react';
import { connect } from 'react-redux';
import { createSelector, createStructuredSelector } from 'reselect';
import PropTypes from 'prop-types';
import isEmpty from 'lodash/isEmpty';

import { selectors } from 'reducers';
import { actions as backendActions } from 'reducers/network/backend';
import { actions as graphActions } from 'reducers/network/graph';
import { actions as pageActions } from 'reducers/network/page';
import { actions as wizardActions } from 'reducers/network/wizard';
import { actions as deploymentActions } from 'reducers/deployments';
import NetworkGraph from 'Components/NetworkGraph';
import NoResultsMessage from 'Components/NoResultsMessage';
import { filterModes } from 'constants/networkFilterModes';
import entityTypes from 'constants/entityTypes';
import wizardStages from '../Wizard/wizardStages';
import Filters from './Overlays/Filters';
import Legend from './Overlays/Legend';

class Graph extends Component {
    static propTypes = {
        wizardOpen: PropTypes.bool.isRequired,
        filterState: PropTypes.number.isRequired,
        isSimulatorOn: PropTypes.bool.isRequired,

        networkNodeMap: PropTypes.shape({}).isRequired,
        networkEdgeMap: PropTypes.shape({}),

        networkPolicyGraphState: PropTypes.string.isRequired,
        networkFlowGraphUpdateKey: PropTypes.number.isRequired,
        networkFlowGraphState: PropTypes.string.isRequired,
        setSelectedNode: PropTypes.func.isRequired,
        setSelectedNamespace: PropTypes.func.isRequired,
        fetchDeployment: PropTypes.func.isRequired,
        fetchNetworkPolicies: PropTypes.func.isRequired,

        setWizardStage: PropTypes.func.isRequired,
        openWizard: PropTypes.func.isRequired,
        closeWizard: PropTypes.func.isRequired,

        clusters: PropTypes.arrayOf(PropTypes.object).isRequired,
        selectedClusterId: PropTypes.string,

        isLoading: PropTypes.bool.isRequired,
    };

    static defaultProps = {
        networkEdgeMap: null,
        selectedClusterId: '',
    };

    shouldComponentUpdate(nextProps) {
        const {
            networkFlowGraphUpdateKey,
            filterState,
            isLoading,
            wizardOpen,
            networkEdgeMap,
            networkNodeMap,
            isSimulatorOn,
        } = this.props;
        return (
            !networkEdgeMap ||
            isEmpty(networkNodeMap) ||
            nextProps.networkFlowGraphUpdateKey !== networkFlowGraphUpdateKey ||
            nextProps.filterState !== filterState ||
            nextProps.isLoading !== isLoading ||
            nextProps.wizardOpen !== wizardOpen ||
            nextProps.isSimulatorOn !== isSimulatorOn
        );
    }

    onNamespaceClick = (namespace) => {
        if (this.props.isSimulatorOn) {
            return;
        }
        this.props.setSelectedNamespace(namespace);
        this.props.setWizardStage(wizardStages.namespaceDetails);
        this.props.openWizard();
    };

    onNodeClick = (node) => {
        if (node?.type === entityTypes.CLUSTER || this.props.isSimulatorOn) {
            return;
        }
        this.props.setSelectedNode(node);
        this.props.fetchDeployment(node.deploymentId);
        this.props.fetchNetworkPolicies([...node.policyIds]);
        this.props.setWizardStage(wizardStages.details);
        this.props.openWizard();
    };

    renderGraph = (simulatorOn) => {
        const { networkNodeMap } = this.props;
        // If we have more than 1100 nodes, display a message instead of the graph.
        const nodeLimit = 1100;
        if (Object.keys(networkNodeMap).length > nodeLimit) {
            // hopefully a temporal solution
            return (
                <NoResultsMessage message="There are too many deployments to render on the graph. Please refine your search to a set of namespaces or deployments to display." />
            );
        }
        const { networkFlowGraphUpdateKey, networkEdgeMap } = this.props;
        const { closeWizard, filterState, clusters, selectedClusterId } = this.props;

        const selectedClusterName =
            clusters.find((cluster) => cluster.id === selectedClusterId)?.name || 'Unknown cluster';

        return (
            <NetworkGraph
                updateKey={networkFlowGraphUpdateKey}
                networkEdgeMap={networkEdgeMap}
                networkNodeMap={networkNodeMap}
                onNodeClick={this.onNodeClick}
                onNamespaceClick={this.onNamespaceClick}
                onClickOutside={closeWizard}
                filterState={filterState}
                simulatorOn={simulatorOn}
                selectedClusterName={selectedClusterName}
            />
        );
    };

    render() {
        const { filterState, isSimulatorOn } = this.props;
        const { networkFlowGraphState, networkPolicyGraphState } = this.props;
        // Simulator styling.
        const simulatorMode = isSimulatorOn ? 'simulator-mode' : '';

        // Graph nodes and styling.
        const networkGraphState =
            filterState === filterModes.active ? networkFlowGraphState : networkPolicyGraphState;
        const networkGraphStateClass = networkGraphState === 'ERROR' ? 'error' : 'success';

        // Rendering.
        return (
            <div className={`${simulatorMode} ${networkGraphStateClass} w-full h-full theme-light`}>
                {this.renderGraph(isSimulatorOn)}
                <Filters />
                <Legend />
            </div>
        );
    }
}

const getIsSimulatorOn = createSelector(
    [selectors.getNetworkWizardOpen, selectors.getNetworkWizardStage],
    (wizardOpen, wizardStage) =>
        wizardOpen &&
        (wizardStage === wizardStages.simulator || wizardStage === wizardStages.creator)
);

const mapStateToProps = createStructuredSelector({
    wizardOpen: selectors.getNetworkWizardOpen,
    filterState: selectors.getNetworkGraphFilterMode,
    isSimulatorOn: getIsSimulatorOn,

    networkNodeMap: selectors.getNetworkNodeMap,
    networkEdgeMap: selectors.getNetworkEdgeMap,

    networkPolicyGraphState: selectors.getNetworkPolicyGraphState,
    networkFlowGraphUpdateKey: selectors.getNetworkFlowGraphUpdateKey,
    networkFlowGraphState: selectors.getNetworkFlowGraphState,

    clusters: selectors.getClusters,
    selectedClusterId: selectors.getSelectedNetworkClusterId,

    isLoading: selectors.getNetworkGraphLoading,
});

const mapDispatchToProps = {
    setSelectedNode: graphActions.setSelectedNode,
    setSelectedNamespace: graphActions.setSelectedNamespace,
    fetchDeployment: deploymentActions.fetchDeployment.request,
    fetchNetworkPolicies: backendActions.fetchNetworkPolicies.request,

    openWizard: pageActions.openNetworkWizard,
    setWizardStage: wizardActions.setNetworkWizardStage,
    setNetworkGraphRef: graphActions.setNetworkGraphRef,
    setNetworkGraphLoading: graphActions.setNetworkGraphLoading,
    closeWizard: pageActions.closeNetworkWizard,
};

export default connect(mapStateToProps, mapDispatchToProps)(Graph);
