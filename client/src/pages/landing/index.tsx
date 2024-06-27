import { CredentialResponse, GoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

import "./index.css";
import SudokuService from "../../services/sudoku";

interface GoogleJwtDecoded {
    email: string;
}

function LandingPage() {
    const navigate = useNavigate();

    const handleLoggedIn = async (credentialResponse: CredentialResponse) => {
        if (!credentialResponse.credential) {
            return "error";
        }

        const userGoogleData = jwtDecode(credentialResponse.credential) as GoogleJwtDecoded;

        const foundUser = await SudokuService.findUserByEmail(userGoogleData.email);

        if (!foundUser) {
            navigate("/register");

            return;
        }

        console.log("userGoogleData", userGoogleData);
        console.log("foundUser", foundUser);

        navigate("/home");
    };

    return (
        <div className="login-container">
            <div className="logo-container">
                <h1 className="logo">
                    Sudoku <span>League</span>
                </h1>
            </div>
            <GoogleLogin
                onSuccess={async (credentialResponse) => {
                    await handleLoggedIn(credentialResponse);
                }}
                onError={() => {
                    console.log("Login Failed");
                }}
            />
        </div>
    );
}

export default LandingPage;
