import React from "react";
import ReactDOM from "react-dom/client";

import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { GoogleOAuthProvider } from "@react-oauth/google";

import reportWebVitals from "./reportWebVitals";
import "./index.css";

import LandingPage from "./pages/landing";
import HomePage from "./pages/home";
import ErrorPage from "./pages/error";

const router = createBrowserRouter([
    {
        path: "/",
        element: <LandingPage />,
        errorElement: <ErrorPage />,
    },
    {
        path: "/home",
        element: <HomePage />,
        errorElement: <ErrorPage />,
    },
]);

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement);
root.render(
    <GoogleOAuthProvider clientId="">
        <React.StrictMode>
            <RouterProvider router={router}></RouterProvider>
        </React.StrictMode>
    </GoogleOAuthProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
