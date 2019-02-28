import React, { Component } from 'react';
import '../App.css';
import ListUsers from './ListUsers';

class UserMain extends Component {
  constructor(props) {
    super(props)
    const isDev = process.env.NODE_ENV !== 'production';
    const port = !isDev ? (process.env.PORT !== undefined ? process.env.PORT : 8081) : 8081;
    const url = (isDev ? "http://localhost:"+ port : "https://goreactcrud.herokuapp.com") + "/api";
    this.state = {
        users : new Set([]),
        isLoading: true,
        apiUrlBase : url,
        getUrl : (endpointKey) => this.state.apiUrlBase + "/" + endpointKey
    }
    console.log('React got port as', this.state.apiUrlBase)
    this.addUser = this.addUser.bind(this);
    this.deleteUser = this.deleteUser.bind(this);
  }

  componentDidMount() {
    document.title = "Go using React UI";
    this.apiCheck().then(() => {
      const timer = setTimeout(() => {
        this.getUsers()
        clearTimeout(timer);
    },2000)
    });
  }

  setAppState = async options => await this.setState({...this.state, ...options});

  async apiCheck() {
    return fetch(this.state.getUrl("healthcheck"), { 
      method: "GET"
    , headers: {
        Accept: "application/json",
      }
    })
    .then(response => response.json()) 
    .then(check => console.log("API health check:", check))
    .catch(error => console.log("Error occured during api health check: ", error))
  }

  async getUsers() {
    await this.setAppState({isLoading:true})
    return fetch(this.state.getUrl("users"), { 
      method: "GET"
    , headers: {
        Accept: "application/json",
      }
    })
    .then(response => response.json()) 
    .then(users => { this.setAppState({isLoading: false , users: new Set((users !== null ? [...users] : []))}) })
    .catch(error => console.log("Error occured while getting user list: ", error))
  }

  async addUser(user) {
    const { Name, Email } = user;
    const appendUrl = "/" + Name + "/" + Email
    return fetch(this.state.getUrl("adduser" + appendUrl), { 
      method: "POST"
    })
    .then(response => { console.log('add user response:', response); this.getUsers(); return true; }) 
    .catch(error => { console.log('Error occured while adding user: ', user, error); return false; });
  }

  async deleteUser(user) {
    const { Name, Email } = user;
    const appendUrl = "/" + Name + "/" + Email
    return fetch(this.state.getUrl("deleteuser" + appendUrl), { 
      method: "DELETE"
    })
    .then(res => { console.log('delete user response:', res); this.getUsers(); return res; }) 
    .catch(error => console.log('Error occured while deleting user: ', user, error));
  }

  render() {
    return (
        <div className="UserMain">
          <div>
            <h1>Users</h1>
          </div>
          <div>
            <ListUsers 
              isLoading={this.state.isLoading}
              users={this.state.users} 
              addUser={(user) => this.addUser(user)} 
              deleteUser={(user) => this.deleteUser(user)}/>
          </div>
      </div>
    );
  }
}

export default UserMain;