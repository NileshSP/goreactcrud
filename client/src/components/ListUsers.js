import React, { Component } from 'react';
import '../App.css';

class ListUsers extends Component {
  constructor(props) {
    super(props);
    this.state = {
      displayMessage: 'loading data . . .'
    }
    this.addUserItem = this.addUserItem.bind(this);
    this.addTextValueChange = this.addTextValueChange.bind(this);
    this.userName = React.createRef();
    this.userEmail = React.createRef();
    this.btnAdd = React.createRef();
  }

  componentDidMount() {
    this.addTextValueChange();
    if(this.props.isLoading) {
      const timer = setTimeout(() => {
          const message = 'data api currently unavailable'
          this.setState({ displayMessage: message });
          clearTimeout(timer);
      },30000)
    } 
  }

  async addTextValueChange() {
    if(this.userName.current.validity.valid && this.userEmail.current.validity.valid) {
      this.btnAdd.current.disabled = false;
    }
    else {
      this.btnAdd.current.disabled = true;
    }
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
    this.addTextValueChange();
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
                    <div><input id="userName" ref={this.userName} type="text" placeholder="Name..." onChange={e => this.addTextValueChange()} pattern=".*\S+.*" required/></div>
                    <div><input id="userEmail" ref={this.userEmail} type="email" placeholder="Email..." onChange={e => this.addTextValueChange()} required/></div>
                    <div><input id="btnAdd" ref={this.btnAdd} type="button" value="Add" className="btn btn-outline-primary btn-sm"
                            onClick={(e) => this.addUserItem() } 
                            />
                    </div>
                </div>
            </li>    
        { 
          (this.props.isLoading) 
          ? 
            <li>
              <div className={(this.state.displayMessage.indexOf('loading') > -1 ? "loading" : "ListUserItems")} >{this.state.displayMessage}</div>
            </li>
          :          
            users.map(user => {
              return  <li key={user.Name+user.Email}>
                    <div className="ListUserItems">
                        <div><input type="text" value={user.Name} readOnly /></div>
                        <div><input type="text" value={user.Email} readOnly /></div>
                        <div><input key={user.ID} type="button" value="Delete" className="btn btn-outline-danger btn-sm"
                                onClick={(e) => this.props.deleteUser(user)} /></div>
                    </div>
                </li>
            })
        }
        </ul>
      </div>
    );
  }
}

export default ListUsers;