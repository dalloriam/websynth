import * as React from 'react';
import './App.css';

import ApolloClient from 'apollo-boost';
import { ApolloProvider } from 'react-apollo';

import TitleBar from './Bar';

import Synth from './synth/Synth';


const client = new ApolloClient({
  uri: "http://localhost:8080/synth"
});


class App extends React.Component {
  public render() {
    return (
      <ApolloProvider client={client}>
        <div className="App">
          <TitleBar />
          <Synth/>
        </div>
      </ApolloProvider>
    );
  }
}

export default App;
