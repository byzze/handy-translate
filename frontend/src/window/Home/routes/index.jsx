import { Navigate } from 'react-router-dom';

import Translate from '../pages/Translate';
import Recognize from '../pages/Recognize';
// import History from '../pages/History';
import Hotkey from '../pages/Hotkey';
import About from '../pages/About';

const routes = [
    {
        path: '/translate',
        element: <Translate />,
    },
    {
        path: '/recognize',
        element: <Recognize />,
    },
    {
        path: '/hotkey',
        element: <Hotkey />,
    },
    {
        path: '/about',
        element: <About />,
    },
    {
        path: '/',
        element: <Navigate to='/hotkey' />,
    },
];

export default routes;
