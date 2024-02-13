import Index from './pages/Index.tsx';
import Login from "./pages/login.tsx";
import Dashboard from "./pages/dashboard.tsx";
import Manage from "./pages/manage.tsx";

const routes = [
    {
        path: '/',
        element: <Index/>
    },
    {
        path: '/login',
        element: <Login/>
    },
    {
        path: '/dashboard',
        element: <Dashboard/>
    },
    {
        path: '/manage',
        element: <Manage/>
    },
]

export default routes