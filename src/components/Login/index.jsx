import "./style.css";
import React, { useState } from "react";
import { Auth } from "aws-amplify";

const Login = () => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  
  
  const login = () => {
    Auth.signIn(username, password).then((result) => {
      console.log("logged in")
     }).catch((err) => {
      console.log(err)
     })
  }
  return (
    <section className="hero is-dark is-fullheight">
  <div className="hero-body">
    <div className="container">
      <div className="columns is-centered">
        <div className="column is-5-tablet is-4-desktop is-3-widescreen">
          <div className="box">
            <div className="field">
              <label className="label">Email</label>
              <div className="control has-icons-left">
                <input type="email"  className="input"  onChange={(evt)=>setUsername(evt.target.value)}/>
                <span className="icon is-small is-left">
                  <i className="fa fa-envelope"></i>
                </span>
              </div>
            </div>
            <div className="field">
              <label className="label">Password</label>
              <div className="control has-icons-left">
                <input type="password" placeholder="*******" className="input"  onChange={(evt)=>setPassword(evt.target.value)}/>
                <span className="icon is-small is-left">
                  <i className="fa fa-lock"></i>
                </span>
              </div>
            </div>
            <div className="field">
              <button className="button is-success" onClick={login}>
                Login
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
  );
};

export default Login;
