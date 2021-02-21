import Redis from 'ioredis';
import { hasNumber } from './hasNumber';
import { handleError } from './handleError';

const redis = new Redis({ keyPrefix: 'user:' });

/**
 * Checks if user has proper security clearance
 * @param userID
 * @param requiredClearanceLevel
 */
export const hasSecurityClearance = async (
  userID: string,
  requiredClearanceLevel: number,
): Promise<boolean> => {
  const response = await redis.get(userID).catch((err) => {
    handleError(err);
  });

  if (typeof response === 'string' && hasNumber(response)) {
    const userClearanceLevel = parseInt(response);

    // if clearance level is greater than or equal to required level, return true
    if (userClearanceLevel >= requiredClearanceLevel)
      return true;
  }

  return false;
};
