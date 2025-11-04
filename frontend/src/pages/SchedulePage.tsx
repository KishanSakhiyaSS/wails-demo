import React, { useState, useEffect } from 'react';
import { Schedule, ScheduleFormData } from '../types/schedule';
import { 
  addSchedule, 
  listSchedules, 
  updateSchedule, 
  deleteSchedule, 
  toggleSchedule 
} from '../services/scheduleService';
import ScheduleForm from '../components/ScheduleForm';
import ScheduleList from '../components/ScheduleList';

const SchedulePage: React.FC = () => {
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [editingSchedule, setEditingSchedule] = useState<Schedule | null>(null);

  // Load schedules on component mount
  useEffect(() => {
    loadSchedules();
  }, []);

  const loadSchedules = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await listSchedules();
      
      // Ensure data is always an array
      setSchedules(Array.isArray(data) ? data : []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load schedules');
      console.error('Error loading schedules:', err);
      // Set empty array on error to prevent further issues
      setSchedules([]);
    } finally {
      setLoading(false);
    }
  };

  const handleAddSchedule = async (formData: ScheduleFormData) => {
    try {
      await addSchedule(formData);
      await loadSchedules();
      setShowForm(false);
      showToast('Schedule added successfully!', 'success');
    } catch (err) {
      showToast(err instanceof Error ? err.message : 'Failed to add schedule', 'error');
      console.error('Error adding schedule:', err);
    }
  };

  const handleUpdateSchedule = async (schedule: Schedule) => {
    try {
      await updateSchedule(schedule);
      await loadSchedules();
      setEditingSchedule(null);
      showToast('Schedule updated successfully!', 'success');
    } catch (err) {
      showToast(err instanceof Error ? err.message : 'Failed to update schedule', 'error');
      console.error('Error updating schedule:', err);
    }
  };

  const handleDeleteSchedule = async (id: number) => {
    if (!confirm('Are you sure you want to delete this schedule?')) {
      return;
    }

    try {
      await deleteSchedule(id);
      await loadSchedules();
      showToast('Schedule deleted successfully!', 'success');
    } catch (err) {
      showToast(err instanceof Error ? err.message : 'Failed to delete schedule', 'error');
      console.error('Error deleting schedule:', err);
    }
  };

  const handleToggleSchedule = async (id: number, enabled: boolean) => {
    try {
      await toggleSchedule(id, enabled);
      await loadSchedules();
      showToast(`Schedule ${enabled ? 'enabled' : 'disabled'} successfully!`, 'success');
    } catch (err) {
      showToast(err instanceof Error ? err.message : 'Failed to toggle schedule', 'error');
      console.error('Error toggling schedule:', err);
    }
  };

  const showToast = (message: string, type: 'success' | 'error') => {
    // Simple toast notification - you can enhance this with a proper toast library
    const toast = document.createElement('div');
    toast.className = `fixed top-4 right-4 px-6 py-3 rounded-lg shadow-lg z-50 ${
      type === 'success' ? 'bg-green-500 text-white' : 'bg-red-500 text-white'
    }`;
    toast.textContent = message;
    document.body.appendChild(toast);

    setTimeout(() => {
      document.body.removeChild(toast);
    }, 3000);
  };

  const handleEditSchedule = (schedule: Schedule) => {
    setEditingSchedule(schedule);
    setShowForm(true);
  };

  const handleFormSubmit = async (data: Schedule | ScheduleFormData) => {
    if (editingSchedule) {
      // Update existing schedule
      await handleUpdateSchedule(data as Schedule);
    } else {
      // Add new schedule
      await handleAddSchedule(data as ScheduleFormData);
    }
  };

  const handleCancelForm = () => {
    setShowForm(false);
    setEditingSchedule(null);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-screen bg-gray-900">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-500 mx-auto mb-4"></div>
          <p className="text-gray-400 text-xl">Loading schedules...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-900 p-6">
      <div className="max-w-6xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-white mb-2">App Scheduler</h1>
          <p className="text-gray-400">
            Schedule automatic start and stop times for your application
          </p>
        </div>

        {/* Error Display */}
        {error && (
          <div className="mb-6 p-4 bg-red-900 border border-red-700 rounded-lg">
            <p className="text-red-200">{error}</p>
            <button
              onClick={loadSchedules}
              className="mt-2 px-4 py-2 bg-red-700 text-white rounded hover:bg-red-600 transition-colors"
            >
              Retry
            </button>
          </div>
        )}

        {/* Action Buttons */}
        <div className="mb-6 flex gap-4">
          <button
            onClick={() => setShowForm(true)}
            className="px-6 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors flex items-center gap-2"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Add New Schedule
          </button>
          
          <button
            onClick={loadSchedules}
            className="px-6 py-3 bg-gray-700 text-white rounded-lg hover:bg-gray-600 transition-colors flex items-center gap-2"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            Refresh
          </button>
        </div>

        {/* Schedule Form Modal */}
        {showForm && (
          <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div className="bg-gray-800 rounded-lg p-6 w-full max-w-md mx-4">
              <ScheduleForm
                schedule={editingSchedule}
                onSubmit={handleFormSubmit}
                onCancel={handleCancelForm}
              />
            </div>
          </div>
        )}

        {/* Schedule List */}
        <ScheduleList
          schedules={schedules}
          onEdit={handleEditSchedule}
          onDelete={handleDeleteSchedule}
          onToggle={handleToggleSchedule}
        />
      </div>
    </div>
  );
};

export default SchedulePage;
