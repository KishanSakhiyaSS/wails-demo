export interface Schedule {
  id: number;
  title: string;
  start_time: string;
  end_time: string;
  repeat_pattern: 'once' | 'daily' | 'weekly';
  enabled: boolean;
  created_at: string;
}

export interface ScheduleFormData {
  title: string;
  start_time: string;
  end_time: string;
  repeat_pattern: 'once' | 'daily' | 'weekly';
  enabled: boolean;
}
