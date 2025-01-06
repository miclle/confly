package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetInstanceConfigSetArgs struct {
	GroupName     string   `uri:"group_name"`
	AppName       string   `uri:"app_name"`
	ConfigSetName string   `uri:"configset_name"`
	ClusterName   string   `query:"clusterName"`
	ServiceName   string   `query:"serviceName"`
	PublishID     string   `query:"curPublishID"`
	IP            string   `query:"ip"`
	Hostname      string   `query:"hostname"`
	Tags          []string `query:"tag"`
}

func (ctrl *Handler) GetConfiguration(c *fox.Context, args *GetInstanceConfigSetArgs) (*models.Configuration, error) {
	c.Logger.Debugf("get configuration args: %+v", args)

	cCp := c.Copy()
	go func() {
		_, err := ctrl.manager.UpsertInstance(cCp, &params.UpsertInstance{
			GroupName:     args.GroupName,
			AppName:       args.AppName,
			ConfigSetName: args.ConfigSetName,
			ClusterName:   args.ClusterName,
			Hostname:      args.Hostname,
			ServiceName:   args.ServiceName,
			IP:            args.IP,
			Tags:          args.Tags,
			PublishID:     args.PublishID,
		})
		if err != nil {
			c.Logger.Errorf("upsert instance failed, error: %+v", err)
		}
	}()

	configSet, err := ctrl.manager.GetConfiguration(c, &params.GetConfiguration{
		GroupName:     args.GroupName,
		AppName:       args.AppName,
		ConfigSetName: args.ConfigSetName,
		ClusterName:   args.ClusterName,
		Hostname:      args.Hostname,
		IP:            args.IP,
		Tags:          args.Tags,
		PublishID:     args.PublishID,
	})

	if err != nil {
		c.Logger.Errorf("get configuration failed, error: %+v", err)
		return nil, err
	}
	return configSet, nil
}

func (ctrl *Handler) ReportInstance(c *fox.Context, args *GetInstanceConfigSetArgs) error {
	c.Logger.Debugf("report instance args: %+v", args)

	_, err := ctrl.manager.UpsertInstance(c, &params.UpsertInstance{
		GroupName:     args.GroupName,
		AppName:       args.AppName,
		ConfigSetName: args.ConfigSetName,
		ClusterName:   args.ClusterName,
		Hostname:      args.Hostname,
		ServiceName:   args.ServiceName,
		IP:            args.IP,
		Tags:          args.Tags,
		PublishID:     args.PublishID,
	})
	if err != nil {
		c.Logger.Errorf("upsert instance failed, error: %+v", err)
		return err
	}
	return nil
}
