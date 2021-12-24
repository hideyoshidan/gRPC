<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Sample;

/**
 * interface
 */
class SampleClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * rpc function arg is for request, return is response
     * @param \Sample\SampleRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Sample(\Sample\SampleRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/sample.Sample/Sample',
        $argument,
        ['\Sample\SampleResponse', 'decode'],
        $metadata, $options);
    }

}
