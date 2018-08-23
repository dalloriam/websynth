import gql from 'graphql-tag';
import * as React from 'react';
import { Query } from 'react-apollo';

import MixChanMdl from "../../../model/MixerChannel";

import { MixerChannel } from './Channel';


export class Channels extends React.Component {

    public render() {
        return (
            <Query query={gql`
                {
                Mixer {
                    Channels {
                        Pan {
                            Value
                        }
                        Volume {
                            Value
                        }
                    }
                }
                }
            `}>
            {({ loading, error, data}) => {
                if (loading) {
                    return <p>Loading...</p>
                }
                if (error) {
                    console.log(error);
                    return <p>Error :( </p>
                }

                return data.Mixer.Channels.map((c: MixChanMdl, i: number) => (
                  <div key={i}>
                    <MixerChannel channel={c} Idx={i} />
                  </div>
                ))
            }}
            </Query>
        );
    }
}