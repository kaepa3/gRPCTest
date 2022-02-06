# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: pancake.proto for package 'pancake.baker'

require 'grpc'
require 'pancake_pb'

module Pancake
  module Baker
    module PancakeBakerService
      # class
      class Service
        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'pancake.baker.PancakeBakerService'

        rpc :Bake, ::Pancake::Baker::BakeRequest, ::Pancake::Baker::BakeResponce
        rpc :Report, ::Pancake::Baker::ReportRequest, ::Pancake::Baker::ReportResponce
      end

      Stub = Service.rpc_stub_class
    end
  end
end
