import React, { Component } from 'react';
import '../App.css';

class ListUsers extends Component {
  constructor(props) {
    super(props);
    this.addUserItem = this.addUserItem.bind(this);
    this.userName = React.createRef();
    this.userEmail = React.createRef();
  }

  async addUserItem() {
    const response = await this.props.addUser({
      Name: this.userName.current.value,
      Email: this.userEmail.current.value
    });
    if(response) { 
      this.userName.current.value = '';
      this.userEmail.current.value = '';
    } 
  }

  render() {
    const users = [...this.props.users];  
    return (
      <div className={["ListUsersMain", "rounded"].join(' ')} >
        <ul>
            <li>
                <div className="ListUserHeader" >
                    <div>User</div>
                    <div>Email</div>
                    <div>Action(option)</div>
                </div>
            </li>    
            <li>
                <div className="ListUserAdd" >
                    <div><input id="userName" ref={this.userName} type="text" placeholder="Name..." /></div>
                    <div><input id="userEmail" ref={this.userEmail} type="text" placeholder="Email..." /></div>
                    <div><input id="btnAdd" type="button" value="Add" className="btn btn-outline-primary btn-sm"
                            onClick={(e) => this.addUserItem() } 
                            />
                    </div>
                </div>
            </li>    
        {
            users.map(user => 
                <li key={user.Name+user.Email}>
                    <div className="ListUserItems">
                        <div><input type="text" value={user.Name} readOnly /></div>
                        <div><input type="text" value={user.Email} readOnly /></div>
                        <div><input key={user.ID} type="button" value="Delete" className="btn btn-outline-danger btn-sm"
                                onClick={(e) => this.props.deleteUser(user)} /></div>
                    </div>
                </li>
            )
        }
        </ul>
      </div>
    );
  }
}

export default ListUsers;