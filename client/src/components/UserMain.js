import React, { Component } from 'react';
import '../App.css';
import ListUsers from './ListUsers';

class UserMain extends Component {
  constructor(props) {
    super(props)
    let inDev = process.env.NODE_ENV !== 'production';
    let port = !inDev ? (process.env.PORT !== undefined ? process.env.PORT : 8081) : 8081;
    this.state = {
        users : new Set(),
        apiUrlBase : "http://localhost:"+ port,
        getUrl : (endpointKey) => this.state.apiUrlBase + "/" + endpointKey
    }
    console.log('React got port as', this.state.apiUrlBase)
    this.addUser = this.addUser.bind(this);
    this.deleteUser = this.deleteUser.bind(this);
  }

  componentDidMount() {
    this.getUsers();
  }

  async setAppState(options) {
    await this.setState({...this.state, ...options})
  }

  async getUsers() {
    fetch(this.state.getUrl("users"), { 
      method: "GET"
    , headers: {
        Accept: "application/json",
      }
    })
    .then(res => { return res.json() }) 
    .then(users => { this.setAppState({users: new Set([...users])}) })
    .catch(error => console.log("Error occured while getting user list: ", error))
  }

  async addUser(user) {
    const { Name, Email } = user;
    if(Name.trim() !== "" && Email.trim() !== "") {
      const appendUrl = "/" + Name + "/" + Email
      return fetch(this.state.getUrl("adduser" + appendUrl), { 
        method: "POST"
      })
      .then(res => { console.log('add user response:', res); this.getUsers(); return res; }) 
      .catch(error => { return error; })
    }
    else {
      return new Promise((resolve,reject) => resolve)
    }
  }

  async deleteUser(user) {
    const { Name, Email } = user;
    const appendUrl = "/" + Name + "/" + Email
    return fetch(this.state.getUrl("deleteuser" + appendUrl), { 
      method: "DELETE"
    })
    .then(res => { console.log('delete user response:', res); this.getUsers(); return res; }) 
    .catch(error => { return error; })
  }

  render() {
    return (
        <div className="UserMain">
          <div>
            <h1>Users</h1>
          </div>
          <div>
            <ListUsers 
              users={this.state.users} 
              addUser={(user) => this.addUser(user)} 
              deleteUser={(user) => this.deleteUser(user)}/>
          </div>
      </div>
    );
  }
}

export default UserMain;