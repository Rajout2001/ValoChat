import React, { createContext, useContext, useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const AuthContext = createContext({
  user: localStorage.getItem('connectedUser') || '',
  logout: () => {},
});

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const memUser = useMemo(()=>localStorage.getItem('connectedUser'),[localStorage.getItem('connectedUser')])
  const [user, setUser] = useState(localStorage.getItem('connectedUser') || '');
  const navigate = useNavigate();

  useEffect(() => {
    const user = localStorage.getItem('connectedUser');
    if (user) {
      setUser(user);
    }else{
      setUser('');
    }
  }, [memUser]);

  const logout = () => {
    localStorage.removeItem('connectedUser');
    localStorage.removeItem('token');
    setUser('');
    navigate('/');
  };

  return (
    <AuthContext.Provider value={{ user, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
