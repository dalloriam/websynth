import * as React from 'react';

import './Bar.css';

import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';


class TitleBar extends React.Component {

    public render() {
        return (
            <div className="root">
            <AppBar position="static" color="primary">
                <Toolbar>
                    <Typography variant="title" color="inherit">
                        WebSynth
                    </Typography>
                </Toolbar>
            </AppBar>
            </div>
        );
    }

}

export default TitleBar;