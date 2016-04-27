#!/usr/bin/env ruby

require 'rest-client'
require 'base64'
require 'json'

def generate_mac(data)
  puts "Challenge: #{data}"
  mac = `create_mac #{data}`
  puts "Mac: #{mac}"
  mac
rescue => e
  puts 'KO'
  puts 'Something bad happened when calling the "create_mac" bin'
  puts "Error: #{e.message}"
  exit 1
end

def main
  print 'Initializing radio... '
  head_hostname = 'head'
  puts 'OK'

  print 'Sending SYN packet... '
  begin
    ack = RestClient.get('http://' + head_hostname + '/syn')
  rescue RestClient::Exception, SocketError, Timeout::Error, Errno::EINVAL, Errno::ECONNRESET, EOFError => e
    puts 'KO'
    puts 'Problem communicating with the head'
    puts "Error: #{e.message}"
    exit 1
  end
  puts 'OK'
  puts 'ACK received'

  begin
    parsed_ack = JSON.parse(ack)
  rescue JSON::JSONError => e
    puts 'Problem parsing the ACK'
    puts "Error: #{e.message}"
    exit 1
  end

  if parsed_ack['challenge'].nil?
    puts 'Problem parsing the ACK'
    exit 1
  end

  print 'Generating the SYN/ACK... '
  mac = generate_mac(parsed_ack['challenge'])
  begin
    enc_mac = Base64.encode64(mac)
    response = { :response => enc_mac }.to_json
  rescue => e
    puts 'KO'
    puts 'Problem generating the SYN/ACK'
    puts "Error: #{e.message}"
    exit 1
  end
  puts 'OK'

  print 'Sending SYN/ACK packet... '
  begin
    ack = RestClient.post('http://' + head_hostname + '/syn/ack', response, :content_type => :json)
  rescue RestClient::Exception, SocketError, Timeout::Error, Errno::EINVAL, Errno::ECONNRESET, EOFError => e
    puts 'KO'
    puts 'Problem communicating with the head'
    puts "Error: #{e.message}"
    exit 1
  end
  puts 'OK'
  puts 'ACK received'

  begin
    parsed_ack = JSON.parse(ack)
    message = Base64.decode64(parsed_ack['message'])
  rescue  => e
    puts 'Problem parsing the ACK or decoding the server message'
    puts "Error: #{e.message}"
    exit 1
  end

  puts  message
  message
end

main
exit 0