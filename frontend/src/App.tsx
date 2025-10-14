import { useEffect, useState } from "react";
import Dashboard from "./components/Dashboard";
import { getUser } from "./services/userService";
import { 
  getCPUInfo, 
  getGPUInfo, 
  getOSInfo, 
  getMemoryInfo, 
  getDiskInfo, 
  getLocationInfo, 
  getHardwareInfo,
  getUsagePercentages
} from "./services/systemService";
import { User, AllSystemData } from "./types/system";

function App() {
  const [, setUser] = useState<User | null>(null);
  const [systemData, setSystemData] = useState<AllSystemData>({
    cpu: null,
    gpu: null,
    gpus: null,
    os: null,
    location: null,
    memory: null,
    disk: null,
    hardware: null,
    usagePercentages: null
  });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchAllSystemData = async () => {
    try {
      const [cpu, gpuData, os, memory, disk, location, hardware, usagePercentages] = await Promise.all([
        getCPUInfo().catch(err => { console.error("CPU Error:", err); return null; }),
        getGPUInfo().catch(err => { console.error("GPU Error:", err); return null; }),
        getOSInfo().catch(err => { console.error("OS Error:", err); return null; }),
        getMemoryInfo().catch(err => { console.error("Memory Error:", err); return null; }),
        getDiskInfo().catch(err => { console.error("Disk Error:", err); return null; }),
        getLocationInfo().catch(err => { console.error("Location Error:", err); return null; }),
        getHardwareInfo().catch(err => { console.error("Hardware Error:", err); return null; }),
        getUsagePercentages().catch(err => { console.error("Usage Error:", err); return null; })
      ]);

      // Handle GPU data - it comes as an array, we want the first GPU for display
      const gpus = Array.isArray(gpuData) ? gpuData : null;
      const gpu = gpus && gpus.length > 0 ? gpus[0] : null;

      setSystemData({ cpu, gpu, gpus, os, location, memory, disk, hardware, usagePercentages });
    } catch (err) {
      console.error("Error fetching system data:", err);
      throw err;
    }
  };

  // Selective refresh functions
  const refreshCPU = async () => {
    try {
      const cpu = await getCPUInfo();
      setSystemData(prev => ({ ...prev, cpu }));
    } catch (err) {
      console.error("Error refreshing CPU:", err);
    }
  };

  const refreshGPU = async () => {
    try {
      const gpuData = await getGPUInfo();
      const gpus = Array.isArray(gpuData) ? gpuData : null;
      const gpu = gpus && gpus.length > 0 ? gpus[0] : null;
      setSystemData(prev => ({ ...prev, gpu, gpus }));
    } catch (err) {
      console.error("Error refreshing GPU:", err);
    }
  };

  const refreshMemory = async () => {
    try {
      const memory = await getMemoryInfo();
      setSystemData(prev => ({ ...prev, memory }));
    } catch (err) {
      console.error("Error refreshing Memory:", err);
    }
  };

  const refreshDisk = async () => {
    try {
      const disk = await getDiskInfo();
      setSystemData(prev => ({ ...prev, disk }));
    } catch (err) {
      console.error("Error refreshing Disk:", err);
    }
  };

  const refreshSystem = async () => {
    try {
      const os = await getOSInfo();
      setSystemData(prev => ({ ...prev, os }));
    } catch (err) {
      console.error("Error refreshing System:", err);
    }
  };

  const refreshHardware = async () => {
    try {
      const hardware = await getHardwareInfo();
      setSystemData(prev => ({ ...prev, hardware }));
    } catch (err) {
      console.error("Error refreshing Hardware:", err);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const userData = await getUser();
        setUser(userData);
        await fetchAllSystemData();
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
      <div className="flex items-center justify-center h-screen bg-gray-900 text-xl text-gray-400">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-500 mx-auto mb-4"></div>
          <p>Loading System Information...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-screen bg-gray-900 text-xl text-red-400">
        <div className="text-center">
          <div className="text-6xl mb-4">⚠️</div>
          <p>Error: {error}</p>
        </div>
      </div>
    );
  }

  const handleRefreshData = async () => {
    try {
      await fetchAllSystemData();
    } catch (err) {
      console.error("Error refreshing data:", err);
    }
  };

  return (
    <div className="h-screen overflow-hidden bg-gray-900">
      <main className="h-full">
        <Dashboard 
          systemData={systemData} 
          onRefresh={handleRefreshData}
          onRefreshCPU={refreshCPU}
          onRefreshGPU={refreshGPU}
          onRefreshMemory={refreshMemory}
          onRefreshDisk={refreshDisk}
          onRefreshSystem={refreshSystem}
          onRefreshHardware={refreshHardware}
        />
      </main>
    </div>
  );
}

export default App;
