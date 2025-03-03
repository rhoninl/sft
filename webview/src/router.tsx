import { createBrowserRouter } from "react-router-dom";
import ErrorPage from "./components/error/error";
import Dashboard from "./pages/dashboard/dashboard";

export const router = createBrowserRouter([
    {
        path: "/*",
        element: <Dashboard />,
        errorElement: <ErrorPage />,
    },
]);
