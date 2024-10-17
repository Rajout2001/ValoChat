import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Error from './components/Error.tsx'
import Login from './pages/Login.tsx'
import { AuthProvider } from './context/AuthProvider.tsx';
import Matchs from './pages/Matchs.tsx'
import { WebSocketProvider } from './context/WebSocketProvider.tsx'
import WebSocketConnection from './context/WebSocketConnection.tsx'



ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <WebSocketProvider>
        <AuthProvider >
          <WebSocketConnection>
            <Routes >
              <Route path="/" element={<App />} errorElement={<Error />} /> {/* 👈 Renders at /app/ */}
              <Route path="/login" element={<Login />} errorElement={<Error />} /> {/* 👈 Renders at /login/ */}
              <Route path="/matchs" element={<Matchs />} errorElement={<Error />} /> {/* 👈 Renders at /matchs/ */}
              <Route path='*' element={<Error />} /> {/* 👈 Renders at any other route */}
            </Routes>
          </WebSocketConnection>
        </AuthProvider>
      </WebSocketProvider>
    </BrowserRouter>
  </React.StrictMode>,
)