#!/usr/bin/env ruby

require 'sinatra'
require 'rest-client' 
require 'yaml'
require 'json'

set :bind, '0.0.0.0'
set :port, 80
set :logging, false

if not File.exist?("/opt/config/config.yml")
  abort("Missing /opt/config/config.yml file!")
end

raise "Not able to start. Missing robot identity..." if ENV["DM2_TOKEN"].nil?

# Load config.yml when app starts
config = YAML.load(File.open("/opt/config/config.yml")) 

challenge_ids = Hash.new

get "/*" do
	endpoint = config["games"].find_all {|g| g[1]["get_challenge"] == params[:splat].first}

	raise "Robot's part not found" if (endpoint.empty?)
	
	url = ENV["DM2_ENDPOINT"] + "/challenge/genByGame/" + endpoint[0][1]['name']

 	splat = params.delete('splat')

    response = JSON.parse(RestClient.get url, :user_token => ENV["DM2_TOKEN"])
    challenge_ids[endpoint[0][1]['name']] = response.delete("id")

    return response.to_json
end

post '/*' do
	endpoint = config["games"].find_all {|g| g[1]["solve_challenge"] == params[:splat].first}

	raise "Robot's part not found" if (endpoint.empty?)

	request.body.rewind

	jdata = JSON.parse(request.body.read)
	jdata["id"] = challenge_ids[endpoint[0][1]['name']]
	
	url = ENV["DM2_ENDPOINT"] + "/challenge/result"

	return RestClient.post url, jdata.to_json, :user_token => ENV["DM2_TOKEN"]
end
