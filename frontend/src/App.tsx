import React, { useEffect, useState } from "react";
import Sidebar from "./components/Sidebar";
import Dashboard from "./components/Dashboard";
import Settings from "./components/Settings";
import { getUser } from "./services/userService";
import { getSystemInfo } from "./services/systemService";

function App() {
  const [user, setUser] = useState<any>(null);
  const [system, setSystem] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState<'dashboard' | 'settings'>('dashboard');

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const [userData, systemData] = await Promise.all([
          getUser(),
          getSystemInfo()
        ]);
        setUser(userData);
        setSystem(systemData);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load data");
        console.error("Error fetching data:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return (
      <div style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#f9fafb',
        fontSize: '20px',
        color: '#6b7280'
      }}>
        Loading...
      </div>
    );
  }

  if (error) {
    return (
      <div style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#f9fafb',
        fontSize: '20px',
        color: '#dc2626'
      }}>
        Error: {error}
      </div>
    );
  }

  const handlePageChange = (page: 'dashboard' | 'settings') => {
    setCurrentPage(page);
  };

  const renderCurrentPage = () => {
    switch (currentPage) {
      case 'dashboard':
        return <Dashboard user={user} system={system} />;
      case 'settings':
        return <Settings />;
      default:
        return <Dashboard user={user} system={system} />;
    }
  };

  return (
    <div style={{ display: 'flex', height: '100vh' }}>
      <Sidebar currentPage={currentPage} onPageChange={handlePageChange} />
      {renderCurrentPage()}
    </div>
  );
}

export default App;
