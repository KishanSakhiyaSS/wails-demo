import React from "react";

type Props = { name: string; role: string };

export default function UserCard({ name, role }: Props) {
  return (
    <div style={{
      backgroundColor: 'white',
      borderRadius: '8px',
      boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1)',
      padding: '24px',
      marginBottom: '24px'
    }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
        <div style={{
          width: '48px',
          height: '48px',
          backgroundColor: '#2563eb',
          borderRadius: '50%',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'white',
          fontSize: '20px',
          fontWeight: 'bold'
        }}>
          {name.charAt(0)}
        </div>
        <div>
          <h3 style={{
            fontSize: '20px',
            fontWeight: '600',
            margin: '0 0 4px 0',
            color: '#111827'
          }}>Hello, {name}!</h3>
          <p style={{
            margin: 0,
            color: '#6b7280'
          }}>Role: <span style={{ fontWeight: '500', color: '#2563eb' }}>{role}</span></p>
        </div>
      </div>
    </div>
  );
}
