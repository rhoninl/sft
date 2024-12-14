import { createBrowserRouter } from "react-router-dom";
import ErrorPage from "./component/error/error";
import App from "./App";
import Devices from "./pages/devices/devices";
import Device from "./pages/devices/detail/device";

export const router = createBrowserRouter([
    {
        path: "/*",
        element: <App />,
        errorElement: <ErrorPage />,
    },
    {
        path: "/devices",
        element: <Devices />,
    },
    {
        path: "/devices/:id",
        element: <Device />,
    }
]);
