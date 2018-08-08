import gql from 'graphql-tag';
import * as React from 'react';
import { Query } from 'react-apollo';

import MixerChannel from "../../../model/MixerChannel";


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

                return data.Mixer.Channels.map((c: MixerChannel, i: number) => (
                  <div key={i}>
                    <p>{`Pan: ${c.Pan.Value}, Volume: ${c.Volume.Value}`}</p>
                  </div>
                ))
            }}
            </Query>
        );
    }
}