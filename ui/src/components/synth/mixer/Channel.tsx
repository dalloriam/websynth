import gql from 'graphql-tag';
import * as React from 'react';

import './Channel.css';

import Knob from '../../ui/Knob';

import { Mutation } from 'react-apollo';

import MixerChan from '../../../model/MixerChannel'


interface IChannelProps {
    Idx: number;
    channel: MixerChan;
}

interface IChannelState {
    Idx: number;
    Pan: number;
    Volume: number;
}

const VOLUME_MUTATION = gql`
    mutation ChangeVolume($idx: Int!, $vol: Float!) {
        Mixer {
            Channel(Idx: $idx) {
                Volume {
                    Set(Value: $vol)
                }
            }
        }
    }
`;

const PAN_MUTATION = gql`
    mutation ChangePan($idx: Int!, $pan: Float!) {
        Mixer {
            Channel(Idx: $idx) {
                Pan {
                    Set(Value: $pan)
                }
            }
        }
    }
`


export class MixerChannel extends React.Component<IChannelProps, IChannelState> {

    constructor(props: IChannelProps) {
        super(props);

        this.state = {
            Idx: props.Idx,
            Pan: props.channel.Pan.Value,
            Volume: props.channel.Volume.Value
        };

        this.onVolumeChanged = this.onVolumeChanged.bind(this);
        this.onPanChanged = this.onPanChanged.bind(this);
    }

    public render() {
        return (
            <div className="mixerchannel">
                <Mutation mutation={VOLUME_MUTATION}>
                    {(changeVolume: any) => (
                        <div className="knob">
                            <Knob onChange={this.onVolumeChanged(changeVolume)} min={0} max={1} value={this.state.Volume} step={0.01}/>
                            <p>Vol</p>
                        </div>
                    )}
                </Mutation>

                <Mutation mutation={PAN_MUTATION}>
                {(changePan: any) => (
                    <div className="knob">
                        <Knob onChange={this.onPanChanged(changePan)} min={-1} max={1} value={this.state.Pan} step={0.01}/>
                        <p>Pan</p>
                    </div>
                )}
                </Mutation>
            </div>
        );
    }

    private onVolumeChanged(fn: any) {
        return (newVolume: number) => {
            const v = Math.floor(newVolume * 100) / 100;
            this.setState({Volume: v}, () => {
                fn({variables: {idx: this.state.Idx, vol: v}});
            });
        }
    }

    private onPanChanged(fn: any) {
        return (newPan: number) => {
            const p = Math.floor(newPan * 100) / 100;
            this.setState({Pan: p}, () => {
                fn({variables: {idx: this.state.Idx, pan: p}});
            })
        }
    }
}