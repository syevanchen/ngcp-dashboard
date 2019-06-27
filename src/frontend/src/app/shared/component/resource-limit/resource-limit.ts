export const defaultResources: Resources = {
  cpuLimit: 12,
  cpuRequestLimitPercent: 3,
  memoryLimit: 64,
  memoryRequestLimitPercent: 3,
  replicaLimit: 32,
};

export class Resources {
  cpuLimit: number;
  cpuRequestLimitPercent: number;
  memoryLimit: number;
  memoryRequestLimitPercent: number;
  replicaLimit: number;

  constructor() {
    Object.keys(defaultResources).forEach(item => {
      this[item] = defaultResources[item];
    });
  }
}
