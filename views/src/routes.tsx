import Index from './pages/Index.tsx';
import Login from "./pages/login.tsx";

const routes = [
    {
        path: '/',
        element: <Index/>
    },
    {
        path: '/login',
        element: <Login/>
    },
]

export default routes