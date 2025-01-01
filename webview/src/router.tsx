import { createBrowserRouter } from "react-router-dom";
import ErrorPage from "./component/error/error";
import App from "./App";
import Dashboard from "./pages/dashboard/dashboard";
import Device from "./pages/dashboard/detail/device";

export const router = createBrowserRouter([
    {
        path: "/*",
        element: <App />,
        errorElement: <ErrorPage />,
    },
    {
        path: "/devices",
        element: <Dashboard />,
    },
    {
        path: "/devices/:id",
        element: <Device />,
    },
    {
        path: "/telemetryservices",
        element: <Dashboard />,
    },
    {
        path: "/dashboard",
        element: <Dashboard />,
    }
]);
