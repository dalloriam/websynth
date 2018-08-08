import * as React from 'react';

import Mixer from './mixer';

class Synth extends React.Component {

    public render() {
        return (
            <div className="synth">
                <h1>This is the synthesizer</h1>
                <Mixer />
            </div>
        );
    }
}

export default Synth;