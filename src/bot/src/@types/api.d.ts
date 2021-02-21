/**
 * Basic response from the star citizen api
 */
import exp from 'constants';

export interface BaseResponse {
  message: string;
  success: 0 | 1;
  source?: string;
  data: null;
}

/**
 * Response data that the stats endpoint returns
 */
export interface StatsResponseData {
  current_live: string;
  current_ptu: string;
  fans: number;
  fleet: number;
  funds: number;
}

/**
 * Implements StatsResponseData as an api response
 */
export interface StatsResponse extends BaseResponse {
  data: null | StatsResponseData;
}

export interface WikiArticle {
  title: string;
  fullURl: string;
  mainImage: string;
  summary: string;
}

export interface UserResponseOrg {
  image: string;
  name: string;
  rank: string;
  sid: string;
}

export interface UserResponseProfile {
  badge: string;
  badge_image: string;
  display: string;
  enlisted: string;
  fluency: string[];
  handle: string;
  id: string;
  image: string;
  page: {
    title: string;
    url: string;
  };
}

export interface UserResponseData {
  organization: UserResponseOrg;
  profile: UserResponseProfile;
}

export interface UserResponse extends BaseResponse {
  data: UserResponseData | null;
}

export interface OrgFocus {
  image: string;
  name: string;
}

export interface OrgResponseData {
  archetype: string;
  banner: string;
  commitment: string;
  focus: {
    primary: OrgFocus;
    secondary: OrgFocus;
  };
  headline: {
    html: string;
    plaintext: string;
  };
  href: string;
  lang: string;
  logo: string;
  members: number;
  name: string;
  recruiting: boolean;
  roleplay: boolean;
  sid: string;
  url: string;
}

export interface OrgResponse extends BaseResponse {
  data: OrgResponseData | null;
}

export interface ShipResponseData {
  afterburner_speed: string;
  beam: string;
  cargocapacity: string;
  chassis_id: string;
  compiled: {};
  description: string;
  focus: string;
  height: string;
  id: string;
  length: string;
  manufacturer: {};
  manufacturer_id: string;
  mass: string;
  max_crew: string;
  media: [];
  min_crew: string;
  name: string;
  pitch_max: string;
  price: string;
  production_note: string | null;
  production_status: string;
  roll_max: string;
  scm_speed: string;
  size: string;
  time_modified: string;
  'time_modified.unfiltered': string;
  type: string;
  url: string;
  xaxis_acceleration: string;
  yaw_max: string;
  yaxis_acceleration: string;
  zaxis_acceleration: string;
}

export interface ShipResponse extends BaseResponse {
  data: ShipResponseData[] | null;
}
