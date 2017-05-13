import React from 'react';
import axios from 'axios';
import { map } from 'lodash';

class Users extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            users: [],
            usersLoaded: false
        };
    }

    componentDidMount() {
        axios.get('/api/v1/users')
            .then(res => {
                this.setState({
                    users: res.data,
                    usersLoaded: true
                });
            });
    }

    renderLoading() {
        return (
            <p> Loading users... </p>
        )
    }

    renderUsers() {
        console.log(this.state.users);
        return map(this.state.users, (user) => {
            return (
                <li key={ user.id }> { user.id }. { user.display_name }</li>
            )
        })
    }

    render() {
        console.log('heyyyyyy');
        return (
            <ul>
                { this.state.usersLoaded ? this.renderUsers() : this.renderLoading() }
            </ul>
        );
    }

}

export default Users;