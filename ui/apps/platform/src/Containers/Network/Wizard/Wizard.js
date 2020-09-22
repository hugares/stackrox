import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import { selectors } from 'reducers';
import { actions as pageActions } from 'reducers/network/page';
import Details from './Details/Details';
import Creator from './Creator/Creator';
import Simulator from './Simulator/Simulator';
import NamespaceDetails from './NamespaceDetails/NamespaceDetails';
import NodesUpdateSection from '../Graph/Overlays/NodesUpdateSection';
import ZoomButtons from '../Graph/Overlays/ZoomButtons';

function Wizard({ wizardOpen, onClose }) {
    const width = wizardOpen ? 'md:w-2/3 lg:w-2/5 min-w-120' : 'w-0';

    return (
        <div
            className={`${width} h-full absolute right-0 bg-primary-200 shadow-lg theme-light network-panel`}
        >
            <NodesUpdateSection />
            <ZoomButtons />

            <Details onClose={onClose} />
            <Creator onClose={onClose} />
            <Simulator onClose={onClose} />
            <NamespaceDetails onClose={onClose} />
        </div>
    );
}

Wizard.propTypes = {
    wizardOpen: PropTypes.bool.isRequired,
    onClose: PropTypes.func.isRequired,
};

const mapStateToProps = createStructuredSelector({
    wizardOpen: selectors.getNetworkWizardOpen,
});

const mapDispatchToProps = {
    onClose: pageActions.closeNetworkWizard,
};

export default connect(mapStateToProps, mapDispatchToProps)(Wizard);
