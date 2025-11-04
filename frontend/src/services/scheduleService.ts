import { Schedule } from '../types/schedule';
import { database } from '../../wailsjs/go/models';

// Convert Wails Schedule to our Schedule type
const convertWailsSchedule = (wailsSchedule: database.Schedule): Schedule => ({
  id: wailsSchedule.id,
  title: wailsSchedule.title,
  start_time: wailsSchedule.start_time,
  end_time: wailsSchedule.end_time,
  repeat_pattern: wailsSchedule.repeat_pattern as 'once' | 'daily' | 'weekly',
  enabled: wailsSchedule.enabled,
  created_at: wailsSchedule.created_at,
});

// Convert our Schedule to Wails Schedule
const convertToWailsSchedule = (schedule: Schedule): database.Schedule => {
  const wailsSchedule = new database.Schedule();
  wailsSchedule.id = schedule.id;
  wailsSchedule.title = schedule.title;
  wailsSchedule.start_time = schedule.start_time;
  wailsSchedule.end_time = schedule.end_time;
  wailsSchedule.repeat_pattern = schedule.repeat_pattern;
  wailsSchedule.enabled = schedule.enabled;
  wailsSchedule.created_at = schedule.created_at;
  return wailsSchedule;
};

// Wails bindings for scheduler operations
export const addSchedule = async (schedule: Omit<Schedule, 'id' | 'created_at'>): Promise<void> => {
  const { AddSchedule } = await import('../../wailsjs/go/app/App');
  const wailsSchedule = new database.Schedule();
  wailsSchedule.title = schedule.title;
  wailsSchedule.start_time = schedule.start_time;
  wailsSchedule.end_time = schedule.end_time;
  wailsSchedule.repeat_pattern = schedule.repeat_pattern;
  wailsSchedule.enabled = schedule.enabled;
  return AddSchedule(wailsSchedule);
};

export const listSchedules = async (): Promise<Schedule[]> => {
  const { ListSchedules } = await import('../../wailsjs/go/app/App');
  const wailsSchedules = await ListSchedules();
  
  // Handle case where wailsSchedules might be null or undefined
  if (!wailsSchedules || !Array.isArray(wailsSchedules)) {
    return [];
  }
  
  return wailsSchedules.map(convertWailsSchedule);
};

export const updateSchedule = async (schedule: Schedule): Promise<void> => {
  const { UpdateSchedule } = await import('../../wailsjs/go/app/App');
  const wailsSchedule = convertToWailsSchedule(schedule);
  return UpdateSchedule(wailsSchedule);
};

export const deleteSchedule = async (id: number): Promise<void> => {
  const { DeleteSchedule } = await import('../../wailsjs/go/app/App');
  return DeleteSchedule(id);
};

export const toggleSchedule = async (id: number, enabled: boolean): Promise<void> => {
  const { ToggleSchedule } = await import('../../wailsjs/go/app/App');
  return ToggleSchedule(id, enabled);
};

export const syncWithSystem = async (): Promise<void> => {
  const { SyncWithSystem } = await import('../../wailsjs/go/app/App');
  return SyncWithSystem();
};
