export type Methods = {
  GET: 'GET';
  POST: 'POST';
  PUT: 'PUT';
  DELETE: 'DELETE';
  PATCH: 'PATCH';
  OPTIONS: 'OPTIONS';
  HEAD: 'HEAD';
  CONNECT: 'CONNECT';
  TRACE: 'TRACE';
};

export interface IApi {
  api_key?: string;
  api_name: string;
  api_description?: string;
  api_endpoint: string;
  api_method: Methods;
  api_params?: string;
  api_headers: string;
  api_body: string;
  api_response: string | object | Array<object | string>;
  api_response_code: string;
  api_response_description: string;
  api_tags: string;
}

export interface IController {
  c_key?: string;
  c_name: string;
  c_description: string;
  c_apis?: IApi[];
}
