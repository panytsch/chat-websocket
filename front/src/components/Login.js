import React from 'react';
import { connect } from "react-redux";
import { withRouter } from "react-router-dom";
import {login} from "../methods/login";

class Login extends React.Component{
    componentDidMount() {
        if (this.props.data.user.token){
            this.props.history.push('/messages');
        }
    }
    input = {
        name: null,
        password: null
    };

    render() {
        return (
            <div>
                <input type="text" ref={e => this.input.name = e}/>
                <br/>
                <input type="text" ref={e => this.input.password = e}/>
                <br/>
                <button onClick={() => this.props.login(this.input.name, this.input.password)}>OK</button>
            </div>
        )
    }
}
const mapDispatchToProps = dispatch => ({
    login: (name, pass) => dispatch(login(name, pass))
});

const mapStateToProps = state => ({
    data: state.userData
});

export default connect(mapStateToProps, mapDispatchToProps)(
    withRouter(Login)
);