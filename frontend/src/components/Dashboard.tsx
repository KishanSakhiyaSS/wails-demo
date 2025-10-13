import React from "react";
import UserCard from "./UserCard";

interface DashboardProps {
  user: { name: string; role: string } | null;
  system: { OS: string; Arch: string } | null;
}

export default function Dashboard({ user, system }: DashboardProps) {
  const handleRefresh = () => {
    window.location.reload();
  };

  return (
    <div style={{
      flex: 1,
      padding: '32px',
      backgroundColor: '#f9fafb',
      minHeight: '100vh'
    }}>
      <div style={{
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: '24px'
      }}>
        <h1 style={{
          fontSize: '32px',
          fontWeight: 'bold',
          color: '#111827',
          margin: 0
        }}>Dashboard</h1>
        <button
          onClick={handleRefresh}
          style={{
            padding: '8px 16px',
            backgroundColor: '#2563eb',
            color: 'white',
            border: 'none',
            borderRadius: '6px',
            fontWeight: '500',
            cursor: 'pointer',
            fontSize: '14px'
          }}
          onMouseEnter={(e) => e.target.style.backgroundColor = '#1d4ed8'}
          onMouseLeave={(e) => e.target.style.backgroundColor = '#2563eb'}
        >
          Refresh Data
        </button>
      </div>

      {user && <UserCard name={user.name} role={user.role} />}
      
      {system && (
        <div style={{
          backgroundColor: 'white',
          borderRadius: '8px',
          boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1)',
          padding: '24px'
        }}>
          <h3 style={{
            fontSize: '20px',
            fontWeight: '600',
            marginBottom: '16px',
            color: '#111827'
          }}>System Information</h3>
          <div style={{
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))',
            gap: '16px'
          }}>
            <div style={{
              padding: '16px',
              backgroundColor: '#f9fafb',
              borderRadius: '6px'
            }}>
              <p style={{
                fontSize: '14px',
                color: '#6b7280',
                margin: '0 0 4px 0'
              }}>Operating System</p>
              <p style={{
                fontWeight: '600',
                color: '#111827',
                margin: 0
              }}>{system.OS}</p>
            </div>
            <div style={{
              padding: '16px',
              backgroundColor: '#f9fafb',
              borderRadius: '6px'
            }}>
              <p style={{
                fontSize: '14px',
                color: '#6b7280',
                margin: '0 0 4px 0'
              }}>Architecture</p>
              <p style={{
                fontWeight: '600',
                color: '#111827',
                margin: 0
              }}>{system.Arch}</p>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
