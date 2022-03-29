import "./Home.css";

import React, { useEffect, useState } from "react";
import client from "../http-common";


const Home = () => {
    const [games, setGames] = useState([])
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);

    useEffect(()=>{
        client.get('/game').then(
            (result) => {
                setGames(result.data)
                setIsLoaded(true);
            },
            (error) => {
                setError(error.message);
                setIsLoaded(true);
            }
        )
    }, []);
    let content;
    if (!isLoaded){
        content =  <div className="fa-3x has-text-centered" style={{ width: "100%" }}>
        <i className="fas fa-spinner fa-spin"></i>
      </div>
    } else if (error) {
        content =  <div
        data-testid="error-icon"
        className="has-text-centered pt-4 pb-4"
        style={{ width: "100%" }}
      >
        <i className="fa-3x fas fa-exclamation-circle"></i>
        <div>{error}</div>
      </div>
    } else {
        content = <table className="table">
        <thead>
          <tr>
            <th>Name</th>
          <th>Platform</th>
          </tr>
        </thead>
        <tbody>
            {games.map((game, i)=>{
               return <tr key={i}>
                    <th>{game.display_name}</th>
                    <td>{game.platform}</td>
                </tr>
              })}
        </tbody>
      </table>
    }
 return(
    <section className="has-background-dark section">
      <div className="container">
        <p className="title has-text-light">Games</p>
        <div className="columns is-centered">
          <div className="column is-8">
            <div className="box">
                {content}
            </div>
        </div>
      </div>
    </div>
  </section>
 )
}
export default Home;
