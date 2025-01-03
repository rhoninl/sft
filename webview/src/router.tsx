import { createBrowserRouter } from "react-router-dom";
import ErrorPage from "./component/error/error";
import App from "./App";
import Dashboard from "./pages/dashboard/dashboard";
import Device from "./pages/dashboard/detail/device";
import Settings from "./pages/settings/settings";
export const router = createBrowserRouter([
    {
        path: "/*",
        element: <Dashboard />,
        errorElement: <ErrorPage />,
    },
    {
        path: "/devices/:id",
        element: <Device />,
    },
]);
