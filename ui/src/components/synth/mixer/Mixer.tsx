import * as React from 'react';

import { Channels } from './Channels';


export class Mixer extends React.Component {

    public render() {
        return (
            <div>
                <h2>Mixer</h2>
                <Channels />
            </div>
        );
    }
}