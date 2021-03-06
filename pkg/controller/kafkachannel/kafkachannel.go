package kafkachannel

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	kafkav1alpha1 "knative.dev/eventing-contrib/kafka/channel/pkg/apis/messaging/v1alpha1"
	"knative.dev/eventing-kafka/pkg/controller/constants"
	"knative.dev/eventing-kafka/pkg/controller/util"
	"knative.dev/eventing/pkg/apis/messaging"
)

// Reconcile The KafkaChannel Itself - After Channel Reconciliation (Add MetaData)
func (r *Reconciler) reconcileKafkaChannel(_ context.Context, channel *kafkav1alpha1.KafkaChannel) error {

	// Get Channel Specific Logger
	logger := util.ChannelLogger(r.logger, channel)

	// Reconcile The KafkaChannel's MetaData
	err := r.reconcileMetaData(channel)
	if err != nil {
		logger.Error("Failed To Reconcile KafkaChannel MetaData", zap.Error(err))
		return fmt.Errorf("failed to reconcile kafkachannel metadata")
	} else {
		logger.Info("Successfully Reconciled KafkaChannel MetaData")
		return nil // Success
	}
}

// Reconcile The KafkaChannels MetaData (Annotations, Labels, etc...)
func (r *Reconciler) reconcileMetaData(channel *kafkav1alpha1.KafkaChannel) error {

	// Update The MetaData (Annotations, Labels, etc...)
	annotationsModified := r.reconcileAnnotations(channel)
	labelsModified := r.reconcileLabels(channel)

	// If The KafkaChannel's MetaData Was Modified
	if annotationsModified || labelsModified {

		// Then Persist Changes To Kubernetes
		_, err := r.kafkaClientSet.MessagingV1alpha1().KafkaChannels(channel.Namespace).Update(channel)
		if err != nil {
			r.logger.Error("Failed To Update KafkaChannel MetaData", zap.Error(err))
			return err
		} else {
			r.logger.Info("Successfully Updated KafkaChannel MetaData")
			return nil
		}
	} else {

		// Otherwise Nothing To Do
		r.logger.Info("Successfully Verified KafkaChannel MetaData")
		return nil
	}
}

// Update KafkaChannel Annotations
func (r *Reconciler) reconcileAnnotations(channel *kafkav1alpha1.KafkaChannel) bool {

	// Get The KafkaChannel's Current Annotations
	annotations := channel.Annotations

	// Initialize The Annotations Map If Empty
	if annotations == nil {
		annotations = make(map[string]string)
	}

	// Track Modified Status
	modified := false

	// Add SubscribableDuckVersion Annotation If Missing Or Incorrect
	if annotations[messaging.SubscribableDuckVersionAnnotation] != constants.SubscribableDuckVersionAnnotationV1Alpha1 {
		annotations[messaging.SubscribableDuckVersionAnnotation] = constants.SubscribableDuckVersionAnnotationV1Alpha1
		modified = true
	}

	// Update The Channel's Annotations
	if modified {
		channel.ObjectMeta.Annotations = annotations
	}

	// Return The Modified Status
	return modified
}

// Update KafkaChannel Labels
func (r *Reconciler) reconcileLabels(channel *kafkav1alpha1.KafkaChannel) bool {

	// Get The KafkaChannel's Current Labels
	labels := channel.Labels

	// Initialize The Labels Map If Empty
	if labels == nil {
		labels = make(map[string]string)
	}

	// Track Modified Status
	modified := false

	// Add Kafka Topic Label If Missing
	topicName := util.TopicName(channel)
	if labels[constants.KafkaTopicLabel] != topicName {
		labels[constants.KafkaTopicLabel] = topicName
		modified = true
	}

	// Add Kafka Secret Label If Missing
	secretName := r.kafkaSecretName(channel) // Can Only Be Called AFTER Topic Reconciliation !!!
	if labels[constants.KafkaSecretLabel] != secretName {
		labels[constants.KafkaSecretLabel] = secretName
		modified = true
	}

	// Update The Channel's Labels
	if modified {
		channel.ObjectMeta.Labels = labels
	}

	// Return The Modified Status
	return modified
}
