// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
)

//go:generate mockgen -package=mocks -destination=../mocks/auditmanager.go -source=auditmanager.go AuditmanagerClient
type AuditmanagerClient interface {
	GetAccountStatus(context.Context, *auditmanager.GetAccountStatusInput, ...func(*auditmanager.Options)) (*auditmanager.GetAccountStatusOutput, error)
	GetAssessment(context.Context, *auditmanager.GetAssessmentInput, ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentOutput, error)
	GetAssessmentFramework(context.Context, *auditmanager.GetAssessmentFrameworkInput, ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentFrameworkOutput, error)
	GetAssessmentReportUrl(context.Context, *auditmanager.GetAssessmentReportUrlInput, ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentReportUrlOutput, error)
	GetChangeLogs(context.Context, *auditmanager.GetChangeLogsInput, ...func(*auditmanager.Options)) (*auditmanager.GetChangeLogsOutput, error)
	GetControl(context.Context, *auditmanager.GetControlInput, ...func(*auditmanager.Options)) (*auditmanager.GetControlOutput, error)
	GetDelegations(context.Context, *auditmanager.GetDelegationsInput, ...func(*auditmanager.Options)) (*auditmanager.GetDelegationsOutput, error)
	GetEvidence(context.Context, *auditmanager.GetEvidenceInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceOutput, error)
	GetEvidenceByEvidenceFolder(context.Context, *auditmanager.GetEvidenceByEvidenceFolderInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceByEvidenceFolderOutput, error)
	GetEvidenceFileUploadUrl(context.Context, *auditmanager.GetEvidenceFileUploadUrlInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFileUploadUrlOutput, error)
	GetEvidenceFolder(context.Context, *auditmanager.GetEvidenceFolderInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFolderOutput, error)
	GetEvidenceFoldersByAssessment(context.Context, *auditmanager.GetEvidenceFoldersByAssessmentInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFoldersByAssessmentOutput, error)
	GetEvidenceFoldersByAssessmentControl(context.Context, *auditmanager.GetEvidenceFoldersByAssessmentControlInput, ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, error)
	GetInsights(context.Context, *auditmanager.GetInsightsInput, ...func(*auditmanager.Options)) (*auditmanager.GetInsightsOutput, error)
	GetInsightsByAssessment(context.Context, *auditmanager.GetInsightsByAssessmentInput, ...func(*auditmanager.Options)) (*auditmanager.GetInsightsByAssessmentOutput, error)
	GetOrganizationAdminAccount(context.Context, *auditmanager.GetOrganizationAdminAccountInput, ...func(*auditmanager.Options)) (*auditmanager.GetOrganizationAdminAccountOutput, error)
	GetServicesInScope(context.Context, *auditmanager.GetServicesInScopeInput, ...func(*auditmanager.Options)) (*auditmanager.GetServicesInScopeOutput, error)
	GetSettings(context.Context, *auditmanager.GetSettingsInput, ...func(*auditmanager.Options)) (*auditmanager.GetSettingsOutput, error)
	ListAssessmentControlInsightsByControlDomain(context.Context, *auditmanager.ListAssessmentControlInsightsByControlDomainInput, ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, error)
	ListAssessmentFrameworkShareRequests(context.Context, *auditmanager.ListAssessmentFrameworkShareRequestsInput, ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentFrameworkShareRequestsOutput, error)
	ListAssessmentFrameworks(context.Context, *auditmanager.ListAssessmentFrameworksInput, ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentFrameworksOutput, error)
	ListAssessmentReports(context.Context, *auditmanager.ListAssessmentReportsInput, ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentReportsOutput, error)
	ListAssessments(context.Context, *auditmanager.ListAssessmentsInput, ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentsOutput, error)
	ListControlDomainInsights(context.Context, *auditmanager.ListControlDomainInsightsInput, ...func(*auditmanager.Options)) (*auditmanager.ListControlDomainInsightsOutput, error)
	ListControlDomainInsightsByAssessment(context.Context, *auditmanager.ListControlDomainInsightsByAssessmentInput, ...func(*auditmanager.Options)) (*auditmanager.ListControlDomainInsightsByAssessmentOutput, error)
	ListControlInsightsByControlDomain(context.Context, *auditmanager.ListControlInsightsByControlDomainInput, ...func(*auditmanager.Options)) (*auditmanager.ListControlInsightsByControlDomainOutput, error)
	ListControls(context.Context, *auditmanager.ListControlsInput, ...func(*auditmanager.Options)) (*auditmanager.ListControlsOutput, error)
	ListKeywordsForDataSource(context.Context, *auditmanager.ListKeywordsForDataSourceInput, ...func(*auditmanager.Options)) (*auditmanager.ListKeywordsForDataSourceOutput, error)
	ListNotifications(context.Context, *auditmanager.ListNotificationsInput, ...func(*auditmanager.Options)) (*auditmanager.ListNotificationsOutput, error)
	ListTagsForResource(context.Context, *auditmanager.ListTagsForResourceInput, ...func(*auditmanager.Options)) (*auditmanager.ListTagsForResourceOutput, error)
}
