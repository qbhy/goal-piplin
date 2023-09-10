import Index from './pages/Index.tsx';
import Login from "./pages/login.tsx";
import Dashboard from "./pages/dashboard.tsx";

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
]

export default routes