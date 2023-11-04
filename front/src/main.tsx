import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import App from './App.tsx'
import Ride from './Ride.tsx'
import Drive from './Drive.tsx'
import Register from './Register.tsx'
import './index.css'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <div>not found</div>,
    children: [
      {
        path: "/",
        index: true,
        element: <Ride />,
      },
      {
        path: "ride",
        element: <Ride />,
      },
      {
        path: "/ride/register",
        element: <Register />,
      },
      {
        path: "drive",
        element: <Drive></Drive>,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
