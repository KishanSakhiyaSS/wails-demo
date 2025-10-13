import React from "react";

interface SidebarProps {
  currentPage: 'dashboard' | 'settings';
  onPageChange: (page: 'dashboard' | 'settings') => void;
}

export default function Sidebar({ currentPage, onPageChange }: SidebarProps) {
  return (
    <div style={{
      width: '256px',
      backgroundColor: '#1f2937',
      color: 'white',
      padding: '24px',
      height: '100vh'
    }}>
      <h1 style={{
        fontSize: '20px',
        fontWeight: 'bold',
        marginBottom: '24px'
      }}>Wails Demo</h1>
      <nav>
        <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
          <li 
            style={{
              padding: '8px 12px',
              backgroundColor: currentPage === 'dashboard' ? '#2563eb' : 'transparent',
              borderRadius: '6px',
              marginBottom: '8px',
              cursor: 'pointer'
            }}
            onClick={() => onPageChange('dashboard')}
            onMouseEnter={(e) => {
              if (currentPage !== 'dashboard') {
                e.target.style.backgroundColor = '#374151';
              }
            }}
            onMouseLeave={(e) => {
              if (currentPage !== 'dashboard') {
                e.target.style.backgroundColor = 'transparent';
              }
            }}>
            Dashboard
          </li>
          <li 
            style={{
              padding: '8px 12px',
              backgroundColor: currentPage === 'settings' ? '#2563eb' : 'transparent',
              borderRadius: '6px',
              marginBottom: '8px',
              cursor: 'pointer'
            }}
            onClick={() => onPageChange('settings')}
            onMouseEnter={(e) => {
              if (currentPage !== 'settings') {
                e.target.style.backgroundColor = '#374151';
              }
            }}
            onMouseLeave={(e) => {
              if (currentPage !== 'settings') {
                e.target.style.backgroundColor = 'transparent';
              }
            }}>
            Settings
          </li>
        </ul>
      </nav>
    </div>
  );
}
