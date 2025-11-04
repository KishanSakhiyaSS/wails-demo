import React, { useState, useEffect } from 'react';
import { Schedule, ScheduleFormData } from '../types/schedule';

interface ScheduleFormProps {
  schedule?: Schedule | null;
  onSubmit: (data: ScheduleFormData | Schedule) => Promise<void>;
  onCancel: () => void;
}

const ScheduleForm: React.FC<ScheduleFormProps> = ({ schedule, onSubmit, onCancel }) => {
  const [formData, setFormData] = useState<ScheduleFormData>({
    title: '',
    start_time: '',
    end_time: '',
    repeat_pattern: 'once',
    enabled: true,
  });

  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    if (schedule) {
      setFormData({
        title: schedule.title,
        start_time: schedule.start_time.slice(0, 16), // Convert to datetime-local format
        end_time: schedule.end_time.slice(0, 16),
        repeat_pattern: schedule.repeat_pattern,
        enabled: schedule.enabled,
      });
    }
  }, [schedule]);

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    if (!formData.title.trim()) {
      newErrors.title = 'Title is required';
    }

    if (!formData.start_time) {
      newErrors.start_time = 'Start time is required';
    }

    if (!formData.end_time) {
      newErrors.end_time = 'End time is required';
    }

    if (formData.start_time && formData.end_time) {
      const startDate = new Date(formData.start_time);
      const endDate = new Date(formData.end_time);
      
      // Check if dates are valid
      if (isNaN(startDate.getTime())) {
        newErrors.start_time = 'Invalid start time format';
      }
      if (isNaN(endDate.getTime())) {
        newErrors.end_time = 'Invalid end time format';
      }
      
      // Only check if both dates are valid
      if (!isNaN(startDate.getTime()) && !isNaN(endDate.getTime()) && endDate <= startDate) {
        newErrors.end_time = 'End time must be after start time';
      }
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm() || isSubmitting) {
      return;
    }

    setIsSubmitting(true);
    try {
      if (schedule) {
        // Update existing schedule
        const updatedSchedule: Schedule = {
          ...schedule,
          ...formData,
          start_time: new Date(formData.start_time).toISOString(),
          end_time: new Date(formData.end_time).toISOString(),
        };
        await onSubmit(updatedSchedule);
      } else {
        // Create new schedule - convert time strings to ISO format
        const newScheduleData: ScheduleFormData = {
          ...formData,
          start_time: new Date(formData.start_time).toISOString(),
          end_time: new Date(formData.end_time).toISOString(),
        };
        await onSubmit(newScheduleData);
      }
    } catch (error) {
      console.error('Form submission error:', error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleInputChange = (field: keyof ScheduleFormData, value: string | boolean) => {
    setFormData(prev => ({ ...prev, [field]: value }));
    // Clear error when user starts typing
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: '' }));
    }
  };

  return (
    <div>
      <h2 className="text-2xl font-bold text-white mb-6">
        {schedule ? 'Edit Schedule' : 'Add New Schedule'}
      </h2>

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Title */}
        <div>
          <label htmlFor="title" className="block text-sm font-medium text-gray-300 mb-2">
            Title *
          </label>
          <input
            type="text"
            id="title"
            value={formData.title}
            onChange={(e) => handleInputChange('title', e.target.value)}
            className={`w-full px-4 py-3 bg-gray-700 border rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500 ${
              errors.title ? 'border-red-500' : 'border-gray-600'
            }`}
            placeholder="Enter schedule title"
          />
          {errors.title && <p className="mt-1 text-sm text-red-400">{errors.title}</p>}
        </div>

        {/* Start Time */}
        <div>
          <label htmlFor="start_time" className="block text-sm font-medium text-gray-300 mb-2">
            Start Time *
          </label>
          <input
            type="datetime-local"
            id="start_time"
            value={formData.start_time}
            onChange={(e) => handleInputChange('start_time', e.target.value)}
            className={`w-full px-4 py-3 bg-gray-700 border rounded-lg text-white focus:outline-none focus:ring-2 focus:ring-purple-500 ${
              errors.start_time ? 'border-red-500' : 'border-gray-600'
            }`}
          />
          {errors.start_time && <p className="mt-1 text-sm text-red-400">{errors.start_time}</p>}
        </div>

        {/* End Time */}
        <div>
          <label htmlFor="end_time" className="block text-sm font-medium text-gray-300 mb-2">
            End Time *
          </label>
          <input
            type="datetime-local"
            id="end_time"
            value={formData.end_time}
            onChange={(e) => handleInputChange('end_time', e.target.value)}
            className={`w-full px-4 py-3 bg-gray-700 border rounded-lg text-white focus:outline-none focus:ring-2 focus:ring-purple-500 ${
              errors.end_time ? 'border-red-500' : 'border-gray-600'
            }`}
          />
          {errors.end_time && <p className="mt-1 text-sm text-red-400">{errors.end_time}</p>}
        </div>

        {/* Repeat Pattern */}
        <div>
          <label htmlFor="repeat_pattern" className="block text-sm font-medium text-gray-300 mb-2">
            Repeat Pattern
          </label>
          <select
            id="repeat_pattern"
            value={formData.repeat_pattern}
            onChange={(e) => handleInputChange('repeat_pattern', e.target.value as 'once' | 'daily' | 'weekly')}
            className="w-full px-4 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white focus:outline-none focus:ring-2 focus:ring-purple-500"
          >
            <option value="once">Once</option>
            <option value="daily">Daily</option>
            <option value="weekly">Weekly</option>
          </select>
        </div>

        {/* Enabled Toggle */}
        <div className="flex items-center">
          <input
            type="checkbox"
            id="enabled"
            checked={formData.enabled}
            onChange={(e) => handleInputChange('enabled', e.target.checked)}
            className="w-4 h-4 text-purple-600 bg-gray-700 border-gray-600 rounded focus:ring-purple-500 focus:ring-2"
          />
          <label htmlFor="enabled" className="ml-2 text-sm font-medium text-gray-300">
            Enable this schedule
          </label>
        </div>

        {/* Form Actions */}
        <div className="flex gap-4 pt-4">
          <button
            type="submit"
            disabled={isSubmitting}
            className={`flex-1 px-6 py-3 bg-purple-600 text-white rounded-lg transition-colors font-medium flex items-center justify-center gap-2 ${
              isSubmitting 
                ? 'opacity-50 cursor-not-allowed' 
                : 'hover:bg-purple-700'
            }`}
          >
            {isSubmitting && (
              <svg className="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            )}
            {isSubmitting 
              ? (schedule ? 'Updating...' : 'Creating...') 
              : (schedule ? 'Update Schedule' : 'Create Schedule')
            }
          </button>
          <button
            type="button"
            onClick={onCancel}
            disabled={isSubmitting}
            className={`flex-1 px-6 py-3 bg-gray-600 text-white rounded-lg transition-colors font-medium ${
              isSubmitting 
                ? 'opacity-50 cursor-not-allowed' 
                : 'hover:bg-gray-500'
            }`}
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
};

export default ScheduleForm;
