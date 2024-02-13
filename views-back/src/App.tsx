import {useRoutes} from "react-router-dom";
import routes from "./routes.tsx";

function App() {
    const Router = () => useRoutes(routes)
    return <Router/>
}

export default App
