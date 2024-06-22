import React from "react";
import "./App.css";

import { GoogleLogin } from "@react-oauth/google";
import MatrixRainBackground from "./components/MatrixRainBackground";

function App() {
    return (
        <div className="app">
            <div className="login-container">
                <div className="logo-container">
                    <div className="logo">Sudoku Match</div>
                    <div className="logo-subtitle"></div>
                </div>
                <GoogleLogin
                    onSuccess={(credentialResponse) => {
                        console.log(credentialResponse);
                    }}
                    onError={() => {
                        console.log("Login Failed");
                    }}
                />
            </div>
            <MatrixRainBackground></MatrixRainBackground>
        </div>
    );
}

export default App;
