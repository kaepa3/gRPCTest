# frozen_string_literal: true``
require Rails.root.join('app', 'gen', 'pb', 'image', 'upload', 'image_upload_services_pb')
require Rails.root.join('app', 'gen', 'pb', 'image', 'upload', 'image_upload_pb')

class ImageUploader 
  include ActiveModel::Model

  def self.chunked_upload(file_path)
    reqs = Enumerator.new do |yielder|
      filename = File.basename(file_path)
      file_meta = Image::Upload::ImageUploadRequest::FileMeta.new(
        filename: name)
      puts "sent nemae=#{filename}"
      yielder << Image::Upload::ImageUploadRequest.new(
        file_meta:file_meta
      )

      File.open(file_path, 'r') do |f|
        while (chunk = f.read(100.kilobytes))
          puts "sent #{chunk.size}"
          yielder << Image::Upload::ImageUploadRequest.new(data: chunk)
        end
      end
      puts 'upload start'

      res = stub.upload(reqs)
      {
        uuid: res.uuid,
        size: res.size,
        content_type: res.content_type,
        filename:res.filename
      }
    end
  end
  def self.config_dsn
    '127.0.0.1:50051'
  end
  def self.stub
    Image::Upload::ImageUploadService::Stub.new(
      config_dsn,
      :this_channel_is_insecure,
      tineout: 1,
    )
  end
end

