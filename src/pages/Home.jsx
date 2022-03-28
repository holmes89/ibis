import "./Home.css";
import React, { useState } from "react";
import { Auth } from "aws-amplify";

const Home = () => {
 return(
    <section className="hero is-dark is-fullheight">
    <div className="hero-body">
      <div className="container">
        <p className="title">Games</p>
        <div className="columns is-centered">
          <div className="column is-8">
            <div className="box">
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
 )
}
export default Home;
