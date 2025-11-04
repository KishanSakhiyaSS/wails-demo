import React from 'react';
import { Schedule } from '../types/schedule';

interface ScheduleListProps {
  schedules: Schedule[];
  onEdit: (schedule: Schedule) => void;
  onDelete: (id: number) => void;
  onToggle: (id: number, enabled: boolean) => void;
}

const ScheduleList: React.FC<ScheduleListProps> = ({ schedules, onEdit, onDelete, onToggle }) => {
  // Safety check to ensure schedules is always an array
  const safeSchedules = Array.isArray(schedules) ? schedules : [];
  
  const formatDateTime = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getRepeatPatternLabel = (pattern: string) => {
    switch (pattern) {
      case 'once':
        return 'Once';
      case 'daily':
        return 'Daily';
      case 'weekly':
        return 'Weekly';
      default:
        return pattern;
    }
  };

  const getStatusBadge = (enabled: boolean) => {
    return (
      <span
        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
          enabled
            ? 'bg-green-900 text-green-200'
            : 'bg-gray-700 text-gray-300'
        }`}
      >
        {enabled ? 'Active' : 'Inactive'}
      </span>
    );
  };

  const getNextRunTime = (schedule: Schedule) => {
    const now = new Date();
    const startTime = new Date(schedule.start_time);
    
    if (schedule.repeat_pattern === 'once') {
      if (startTime > now) {
        return `Next run: ${formatDateTime(schedule.start_time)}`;
      } else {
        return 'Completed';
      }
    } else if (schedule.repeat_pattern === 'daily') {
      const today = new Date(now);
      today.setHours(startTime.getHours(), startTime.getMinutes(), 0, 0);
      
      if (today > now) {
        return `Next run: Today at ${startTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })}`;
      } else {
        const tomorrow = new Date(today);
        tomorrow.setDate(tomorrow.getDate() + 1);
        return `Next run: Tomorrow at ${startTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })}`;
      }
    } else if (schedule.repeat_pattern === 'weekly') {
      const daysUntilNext = (startTime.getDay() - now.getDay() + 7) % 7;
      const nextRun = new Date(now);
      nextRun.setDate(nextRun.getDate() + daysUntilNext);
      nextRun.setHours(startTime.getHours(), startTime.getMinutes(), 0, 0);
      
      if (daysUntilNext === 0 && startTime.getHours() > now.getHours()) {
        return `Next run: Today at ${startTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })}`;
      } else {
        return `Next run: ${nextRun.toLocaleDateString('en-US', { weekday: 'long' })} at ${startTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })}`;
      }
    }
    
    return 'Unknown';
  };

  if (safeSchedules.length === 0) {
    return (
      <div className="bg-gray-800 rounded-lg p-8 text-center">
        <div className="text-gray-400 mb-4">
          <svg className="w-16 h-16 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
        </div>
        <h3 className="text-xl font-semibold text-white mb-2">No schedules yet</h3>
        <p className="text-gray-400 mb-4">
          Create your first schedule to automatically start and stop your application.
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {safeSchedules.map((schedule) => (
        <div
          key={schedule.id}
          className="bg-gray-800 rounded-lg p-6 border border-gray-700 hover:border-gray-600 transition-colors"
        >
          <div className="flex items-start justify-between">
            <div className="flex-1">
              <div className="flex items-center gap-3 mb-2">
                <h3 className="text-lg font-semibold text-white">{schedule.title}</h3>
                {getStatusBadge(schedule.enabled)}
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm text-gray-300">
                <div>
                  <span className="font-medium">Start:</span> {formatDateTime(schedule.start_time)}
                </div>
                <div>
                  <span className="font-medium">End:</span> {formatDateTime(schedule.end_time)}
                </div>
                <div>
                  <span className="font-medium">Repeat:</span> {getRepeatPatternLabel(schedule.repeat_pattern)}
                </div>
                <div>
                  <span className="font-medium">Next Run:</span> {getNextRunTime(schedule)}
                </div>
              </div>
            </div>

            <div className="flex items-center gap-2 ml-4">
              {/* Toggle Button */}
              <button
                onClick={() => onToggle(schedule.id, !schedule.enabled)}
                className={`px-3 py-1.5 rounded text-sm font-medium transition-colors ${
                  schedule.enabled
                    ? 'bg-yellow-600 hover:bg-yellow-700 text-white'
                    : 'bg-green-600 hover:bg-green-700 text-white'
                }`}
                title={schedule.enabled ? 'Disable schedule' : 'Enable schedule'}
              >
                {schedule.enabled ? 'Disable' : 'Enable'}
              </button>

              {/* Edit Button */}
              <button
                onClick={() => onEdit(schedule)}
                className="px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white rounded text-sm font-medium transition-colors"
                title="Edit schedule"
              >
                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>

              {/* Delete Button */}
              <button
                onClick={() => onDelete(schedule.id)}
                className="px-3 py-1.5 bg-red-600 hover:bg-red-700 text-white rounded text-sm font-medium transition-colors"
                title="Delete schedule"
              >
                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default ScheduleList;
