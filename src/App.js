import './App.css';
import React, { useEffect, useState } from "react";

import Amplify, { Auth, Hub } from 'aws-amplify';
import Login from './components/Login';

Amplify.configure({
  Auth: {
    userPoolId: 'us-east-2_Ekcud6vxr', //UserPool ID
    region: 'us-east-2  ',
    userPoolWebClientId: '2e21r8j61uje52ju1vsu0ae1uq' //WebClientId
  }
});

const App = () => {
  const [user, setUser] = useState(null);
  
  useEffect(() => {
    Hub.listen("auth", ({ payload: { event, data } }) => {
      debugger;
      switch (event) {
        case "signIn":
        case "cognitoHostedUI":
          getUser().then((userData) => setUser(userData));
          break;
        case "signOut":
          setUser(null);
          break;
        case "signIn_failure":
        case "cognitoHostedUI_failure":
          console.log("Sign in failure", data);
          break;
      }
    });

    getUser().then((userData) => setUser(userData));
  }, []);

  const getUser = async () => {
    return Auth.currentAuthenticatedUser()
      .then((userData) => userData)
      .catch(() => console.log("Not signed in"));
  };
  return (
    <div className="App">
      {user ? (
        <h2>Logged In</h2>
      ) : (
       <Login />
      )}
    </div>
  );
}

export default App;
